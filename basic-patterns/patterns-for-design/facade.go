package main

import "fmt"

// Подсистема 1: Освещение
type Lights struct{}

func (l *Lights) Off() {
	fmt.Println("Свет: выключен")
}

// Подсистема 2: Шторы
type Curtains struct{}

func (c *Curtains) Close() {
	fmt.Println("Шторы: закрыты")
}

// Подсистема 3: Кондиционер
type Thermostat struct{}

func (t *Thermostat) SetTemperature(temp int) {
	fmt.Printf("Кондиционер: Установлена температура %d°C\n", temp)
}

// Подсистема 4: Сигнализация
type Alarm struct{}

func (a *Alarm) Activate() {
	fmt.Println("Сигнализация: активирована")
}

// Фасад: Умный дом
type SmartHomeFacade struct {
	lights     *Lights
	curtains   *Curtains
	thermostat *Thermostat
	alarm      *Alarm
}

// Конструктор фасада
func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lights:     &Lights{},
		curtains:   &Curtains{},
		thermostat: &Thermostat{},
		alarm:      &Alarm{},
	}
}

// Метод для включения режима "Спокойной ночи"
func (s *SmartHomeFacade) GoodNightMode() {
	fmt.Println("Активация режима `Спокойной ночи`...")
	s.lights.Off()
	s.curtains.Close()
	s.thermostat.SetTemperature(20) // Устанавливаем комфортную температуру
	s.alarm.Activate()
	fmt.Println("Режим `Спокойной ночи` активирован!")
}

func main() {
	// Создаём фасад для умного дома
	smartHome := NewSmartHomeFacade()

	// Активируем режим "Спокойной ночи"
	smartHome.GoodNightMode()
}
