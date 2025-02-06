package main

import "fmt"

// Product интерфейс для продуктов
type Product interface {
	GetName() string
}

// ConcreteProductA конкретная реализация продукта A
type ConcreteProductA struct{}

func (p *ConcreteProductA) GetName() string {
	return "ConcreteProductA"
}

// ConcreteProductB конкретная реализация продукта B
type ConcreteProductB struct{}

func (p *ConcreteProductB) GetName() string {
	return "ConcreteProductB"
}

// Factory интерфейс для фабрики
type Factory interface {
	CreateProduct() Product
}

// ConcreteFactoryA конкретная реализация фабрики для продукта A
type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteFactoryB конкретная реализация фабрики для продукта B
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	factories := []Factory{
		&ConcreteFactoryA{},
		&ConcreteFactoryB{},
	}

	for _, factory := range factories {
		product := factory.CreateProduct()
		fmt.Println("Created product:", product.GetName())
	}
}
