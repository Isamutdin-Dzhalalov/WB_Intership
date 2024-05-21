/*
	Паттерн "Фабричный метод" — это порождающий шаблон проектирования, 
	который предоставляет интерфейс для создания объектов в суперклассе, позволяя подклассам изменять тип создаваемых объектов. 

	+:
		- Ослабление зависимости: Класс, использующий фабрику, не зависит от конкретных классов продуктов.
		- Расширяемость: Легко добавить новые продукты, создавая новые классы продуктов и соответствующие фабрики.

	-:
		- Сложность кода: Появляется больше классов и кода, что может сделать систему сложнее.
		Затраты на поддержку: Требуется поддерживать все созданные фабрики и продукты.
*/

package main

import "fmt"

// Интерфейс, который содержит метод, который возвращает строку.
type DBConnection interface {
	Connect() string
}

// Pеализация подключения для MySQL.
type MySQLConnection struct{}

func (m *MySQLConnection) Connect() string {
	return "Connected to MySQL database"
}

// Pеализация подключения для PostgreSQL.
type PostgreSQLConnection struct{}

func (p *PostgreSQLConnection) Connect() string {
	return "Connected to PostgreSQL database"
}

// абстрактный класс фабрики.
type DBFactory interface {
	CreateConnection() DBConnection
}

// конкретная фабрика для MySQL.
type MySQLFactory struct{}

func (m *MySQLFactory) CreateConnection() DBConnection {
	return &MySQLConnection{}
}

// конкретная фабрика для PostgreSQL
type PostgreSQLFactory struct{}

func (p *PostgreSQLFactory) CreateConnection() DBConnection {
	return &PostgreSQLConnection{}
}

/* Создаются подключения к БД, вызывая метод CreateConnection(), для каждой из фабрик.
   Для каждого созданного соединения вызывается метод Connect(), который выводит строку о подключении.*/
func main() {
	var factory DBFactory
	var connection DBConnection

	// Создание подключения к MySQL.
	factory = &MySQLFactory{}
	connection = factory.CreateConnection()
	fmt.Println(connection.Connect())

	// Создание подключения к PostgreSQL.
	factory = &PostgreSQLFactory{}
	connection = factory.CreateConnection()
	fmt.Println(connection.Connect())
}

