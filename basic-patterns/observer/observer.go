package main

import "fmt"

// Observer интерфейс для наблюдателей
type Observer interface {
	Update(message string)
}

// Subject интерфейс для субъекта
type Subject interface {
	RegisterObserver(observer Observer)
	RemoveObserver(observer Observer)
	NotifyObservers(message string)
}

// ConcreteSubject реализация субъекта
type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) RegisterObserver(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) RemoveObserver(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers(message string) {
	for _, observer := range s.observers {
		observer.Update(message)
	}
}

// ConcreteObserver реализация наблюдателя
type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Update(message string) {
	fmt.Printf("%s received message: %s\n", o.name, message)
}

func main() {
	subject := &ConcreteSubject{}

	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}

	subject.RegisterObserver(observer1)
	subject.RegisterObserver(observer2)

	subject.NotifyObservers("Hello, Observers!")

	subject.RemoveObserver(observer1)

	subject.NotifyObservers("Observer1 has been removed.")
}
