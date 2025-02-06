package main

import "fmt"

// Интерфейс, который определяет стратегию оплаты
type PaymentStrategy interface {
	Pay(amount float64)
}

// Стратегия оплаты картой
type CardPayment struct{}

func (c *CardPayment) Pay(amount float64) {
	fmt.Printf("Оплата картой: %.2f рублей\n", amount)
}

// Стратегия оплаты наличными
type CashPayment struct{}

func (c *CashPayment) Pay(amount float64) {
	fmt.Printf("Оплата наличными: %.2f рублей\n", amount)
}

// Контекст, который использует одну из стратегий
type Shop struct {
	paymentStrategy PaymentStrategy
}

func (s *Shop) SetPaymentStrategy(strategy PaymentStrategy) {
	s.paymentStrategy = strategy
}

func (s *Shop) MakePayment(amount float64) {
	s.paymentStrategy.Pay(amount)
}

func main() {
	// Создаем магазин
	shop := &Shop{}

	// Платим картой
	shop.SetPaymentStrategy(&CardPayment{})
	shop.MakePayment(1000.50)

	// Платим наличными
	shop.SetPaymentStrategy(&CashPayment{})
	shop.MakePayment(500.75)
}
