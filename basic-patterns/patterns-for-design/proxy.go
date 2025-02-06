package main

import "fmt"

// Интерфейс для работы с базой данных
type Database interface {
	Connect() string
	Query(query string) string
}

// Реальная база данных, которая выполняет запросы
type RealDatabase struct{}

func (db *RealDatabase) Connect() string {
	return "Подключение к реальной базе данных..."
}

func (db *RealDatabase) Query(query string) string {
	return fmt.Sprintf("Запрос к базе данных: %s", query)
}

// Прокси для базы данных, который проверяет права доступа пользователя
type DatabaseProxy struct {
	realDatabase Database
	userRole     string // Роль пользователя (например, "admin", "user", "guest")
}

func (proxy *DatabaseProxy) Connect() string {
	// Прокси проверяет права доступа
	if proxy.userRole != "admin" {
		return "Ошибка доступа: недостаточно прав для подключения к базе данных."
	}
	// Передаем запрос реальной базе данных
	return proxy.realDatabase.Connect()
}

func (proxy *DatabaseProxy) Query(query string) string {
	// Прокси проверяет права доступа
	if proxy.userRole != "admin" {
		return "Ошибка доступа: недостаточно прав для выполнения запроса."
	}
	// Передаем запрос реальной базе данных
	return proxy.realDatabase.Query(query)
}

func main() {
	// Создаем реальную базу данных
	realDB := &RealDatabase{}

	// Создаем прокси для базы данных с ролью "admin"
	adminProxy := &DatabaseProxy{
		realDatabase: realDB,
		userRole:     "admin", // Этот пользователь имеет доступ
	}

	// Попытка подключиться и выполнить запрос с правами администратора
	fmt.Println(adminProxy.Connect())
	fmt.Println(adminProxy.Query("SELECT * FROM users"))

	// Создаем прокси для базы данных с ролью "guest"
	guestProxy := &DatabaseProxy{
		realDatabase: realDB,
		userRole:     "guest", // У этого пользователя нет доступа
	}

	// Попытка подключиться и выполнить запрос с правами гостя
	fmt.Println(guestProxy.Connect())
	fmt.Println(guestProxy.Query("SELECT * FROM users"))
}
