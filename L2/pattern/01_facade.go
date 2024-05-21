/* Фасад — это структурный паттерн проектирования, который предоставляет простой 
   интерфейс к сложной системе классов, библиотеке или фреймворку.
   Фасад может иметь урезанный интерфейс, не имеющий 100% функциональности, которой можно достичь, 
   используя сложную подсистему напрямую. Но он предоставляет именно те фичи, которые нужны клиенту, и скрывает все остальные.
	
	+:
		- Снижение сложности - предоставляет простой интерфейс, скрывая детали реализации.
		- Удобство использования - клиенты получают доступ к функциональности через единый интерфейс.
		- Инкапсуляция - скрытие деталей - фасад скрывает внутренние детали реализации подсистем,
		  предоставлял только необходимые методы. Повышает безопасность системы.
	-:
		- Дополнительный уровень абстракции - фасад добавляет ещё один уровень абстракции, что может
		  привести к небольшому снижению производительности из-за вызовов дополнительных методов.
		- Ограничение функционала - фасад предоставляет ограниченный набор методов.

	Примеры:
		- Работа с БД. Например упрощение выполнение каких-то операций.
		- Интеграция с внешними API. Фасад может скрыть такие детали, как:
		  Аутентификация, запросы, ответы.
*/

package main

import (
	"fmt"
)

type IgnitionLock struct{}
type Battery struct{}
type Engine struct{}

// Метод структуры.
func (i *IgnitionLock)TurningTheKey() {
	fmt.Println("Ключ повёрнут.")
}

// Метод структуры.
func (e *Engine) LaunchEngine() {
	fmt.Println("Двигатель запущен.")
}

// Метод структуры.
func (b *Battery)CurrentSupply(voltage int) {
	if voltage > 23 {
		fmt.Println("Ток подан на стартер.")
	} else {
		fmt.Println("Аккумулятор разряжен.")
	}
}

// Ф-ция создаёт новый объект и инициализирует все подсистемы.
func NewCar() *Car {
	return &Car {
		ignitionLock: &IgnitionLock{},
		battery: &Battery{},
		engine: &Engine{},
	}
}

// Фасад - cтруктура, которая содержит ссылки на подсистемы.
type Car struct {
	ignitionLock *IgnitionLock
	battery *Battery
	engine *Engine
}

// Метод запускает двигатель проделывая определённые шаги.
func (c *Car) Launch() {
	c.ignitionLock.TurningTheKey()
	c.battery.CurrentSupply(24)
	c.engine.LaunchEngine()
}

func main() {
	car := NewCar()
	car.Launch()
}