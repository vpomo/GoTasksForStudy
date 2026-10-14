package main

import (
	"sync"
	"time"
)

type entry struct {
	value     interface{}
	expiresAt time.Time
}

// TTLCache — потокобезопасный кэш с TTL.
type TTLCache struct {
	mu              sync.RWMutex
	data            map[interface{}]*entry
	cleanupInterval time.Duration
	stop            chan struct{}
	wg              sync.WaitGroup
}

// NewTTLCache создаёт новый кэш.
// cleanupInterval — как часто запускать очистку (например, 10*time.Second).
// Если cleanupInterval <= 0, фоновая очистка не запускается.
func NewTTLCache(cleanupInterval time.Duration) *TTLCache {
	c := &TTLCache{
		data:            make(map[interface{}]*entry),
		cleanupInterval: cleanupInterval,
		stop:            make(chan struct{}),
	}
	if cleanupInterval > 0 {
		c.wg.Add(1)
		go c.cleanupLoop()
	}
	return c
}

// Set добавляет/обновляет значение с заданным TTL.
func (c *TTLCache) Set(key interface{}, value interface{}, ttl time.Duration) {
	expiresAt := time.Now().Add(ttl)

	c.mu.Lock()
	c.data[key] = &entry{
		value:     value,
		expiresAt: expiresAt,
	}
	c.mu.Unlock()
}

// Get возвращает значение и флаг наличия. Если ключ истёк, считается отсутствующим.
func (c *TTLCache) Get(key interface{}) (value interface{}, ok bool) {
	c.mu.RLock()
	e, exists := c.data[key]
	c.mu.RUnlock()

	if !exists {
		return nil, false
	}

	// Проверка TTL вне мьютекса — это нормально: если запись удалится позже,
	// мы просто вернём «не найдено» при следующем запросе.
	if time.Now().After(e.expiresAt) {
		// Удаляем просроченную запись (опционально, чтобы не раздувать кэш)
		c.mu.Lock()
		_, existsInMap := c.data[key]
		if existsInMap && c.data[key].expiresAt.Before(time.Now()) {
			delete(c.data, key)
		}
		c.mu.Unlock()
		return nil, false
	}

	return e.value, true
}

// Delete удаляет ключ.
func (c *TTLCache) Delete(key interface{}) {
	c.mu.Lock()
	delete(c.data, key)
	c.mu.Unlock()
}

// Len возвращает текущее количество элементов (без учёта TTL).
func (c *TTLCache) Len() int {
	c.mu.RLock()
	l := len(c.data)
	c.mu.RUnlock()
	return l
}

// Stop останавливает фоновый цикл очистки. Не блокирует долгоживущие операции.
func (c *TTLCache) Stop() {
	close(c.stop)
	c.wg.Wait()
}

func (c *TTLCache) cleanupLoop() {
	defer c.wg.Done()

	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			c.mu.Lock()
			for k, e := range c.data {
				if now.After(e.expiresAt) {
					delete(c.data, k)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}
