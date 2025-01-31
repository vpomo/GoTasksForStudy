package main

import (
	"fmt"
	"strings"
)

// Filter интерфейс для всех фильтров
type Filter interface {
	Apply(data string) string
}

// BaseFilter структура для базового фильтра
type BaseFilter struct {
	next Filter
}

// SetNext устанавливает следующий фильтр в цепочке
func (f *BaseFilter) SetNext(next Filter) {
	f.next = next
}

// Apply применяет фильтр и передает данные следующему фильтру
func (f *BaseFilter) Apply(data string) string {
	if f.next != nil {
		return f.next.Apply(data)
	}
	return data
}

// ToUpperFilter структура для фильтра, преобразующего строку в верхний регистр
type ToUpperFilter struct {
	BaseFilter
}

// Apply применяет фильтр и передает данные следующему фильтру
func (f *ToUpperFilter) Apply(data string) string {
	data = strings.ToUpper(data)
	return f.BaseFilter.Apply(data)
}

// TrimSpacesFilter структура для фильтра, удаляющего пробелы из строки
type TrimSpacesFilter struct {
	BaseFilter
}

// Apply применяет фильтр и передает данные следующему фильтру
func (f *TrimSpacesFilter) Apply(data string) string {
	data = strings.ReplaceAll(data, " ", "")
	return f.BaseFilter.Apply(data)
}

// AddPrefixFilter структура для фильтра, добавляющего префикс к строке
type AddPrefixFilter struct {
	BaseFilter
	prefix string
}

// NewAddPrefixFilter создает новый фильтр с префиксом
func NewAddPrefixFilter(prefix string) *AddPrefixFilter {
	return &AddPrefixFilter{prefix: prefix}
}

// Apply применяет фильтр и передает данные следующему фильтру
func (f *AddPrefixFilter) Apply(data string) string {
	data = f.prefix + data
	return f.BaseFilter.Apply(data)
}

func main() {
	// Создаем фильтры
	toUpperFilter := &ToUpperFilter{}
	trimSpacesFilter := &TrimSpacesFilter{}
	addPrefixFilter := NewAddPrefixFilter("PREFIX_")

	// Устанавливаем цепочку фильтров
	toUpperFilter.SetNext(trimSpacesFilter)
	trimSpacesFilter.SetNext(addPrefixFilter)

	// Применяем цепочку фильтров к данным
	inputData := "  hello world  "
	result := toUpperFilter.Apply(inputData)

	// Выводим результат
	fmt.Println("Original:", inputData)
	fmt.Println("Filtered:", result)
}

// Original:   hello world
// Filtered: PREFIX_HELLOWORLD
