package main

import (
	"fmt"
)

// ------------------------------------
// Целевой интерфейс, который ожидает клиент
// ------------------------------------
type Printer interface {
	Print(text string)
}

// ------------------------------------
// Современный принтер, реализующий Printer
// ------------------------------------
type ModernPrinter struct{}

func (mp *ModernPrinter) Print(text string) {
	fmt.Println("Modern Printer:", text)
}

// ------------------------------------
// Сторонний (старый) принтер с несовместимым интерфейсом
// ------------------------------------
type LegacyPrinter struct{}

func (lp *LegacyPrinter) PrintLegacy(text []byte) (int, error) {
	n, err := fmt.Println("Legacy Printer:", string(text))
	return n, err
}

// ------------------------------------
// Адаптер для LegacyPrinter, чтобы он соответствовал интерфейсу Printer
// ------------------------------------
type LegacyPrinterAdapter struct {
	legacyPrinter *LegacyPrinter
}

func (adapter *LegacyPrinterAdapter) Print(text string) {
	// Преобразуем string в []byte и вызываем старый метод
	adapter.legacyPrinter.PrintLegacy([]byte(text))
}

// ------------------------------------
// Клиентский код, работающий через интерфейс Printer
// ------------------------------------
func main() {
	printers := []Printer{
		&ModernPrinter{},
		&LegacyPrinterAdapter{legacyPrinter: &LegacyPrinter{}},
	}

	for _, printer := range printers {
		printer.Print("Hello, World!")
	}
}
