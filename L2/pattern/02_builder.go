/*
	Паттерн Builder - это порождающий паттерн проектирования, 
	который используется для создания объектов сложной структуры. 
	Builder позволяет создавать объект последовательным вызовом методов.

	+:
	   - Упрощение создания сложных объектов: Builder позволяет создавать объекты 
	     с множеством параметров без необходимости передачи всех этих параметров через конструктор.
	   - Сокрытие сложности: Builder позволяет сокрыть сложность создания объектов от клиентского кода,
	     что делает код более читаемым и поддерживаемым.
	
	-:
	  - Дополнительный код: Реализация паттерна Builder требует написания дополнительного кода для 
	    создания класса-строителя и методов конфигурации, что может увеличить сложность кода.
	  - Нарушение инкапсуляции: Builder может нарушить инкапсуляцию, если клиентский 
	    код напрямую использует методы конфигурации объекта.

	Практика:
		- Паттерн Builder особенно полезен, когда нужно создавать объекты с 
		  большим количеством необязательных параметров.
		- Builder может быть полезен, когда структура объекта может изменяться 
		  в зависимости от определенных условий или параметров.
*/

package main

import "fmt"

// Продукт, который мы создаем.
type House struct {
    walls   string
    roof    string
    windows string
}

// Интерфейс Строителя, с методами построения и получения готового дома.
type HouseBuilder interface {
    BuildWalls()
    BuildRoof()
    BuildWindows()
    GetHouse() House
}

// Конкретный строитель для деревянного дома.
type WoodenHouseBuilder struct {
    house House
}

func (b *WoodenHouseBuilder) BuildWalls() {
    b.house.walls = "Wooden Walls"
}

func (b *WoodenHouseBuilder) BuildRoof() {
    b.house.roof = "Wooden Roof"
}

func (b *WoodenHouseBuilder) BuildWindows() {
    b.house.windows = "Wooden Windows"
}

func (b *WoodenHouseBuilder) GetHouse() House {
    return b.house
}

// Управление процессом строительства.
type Director struct {
    builder HouseBuilder
}

// Вызываем методы.
func (d *Director) Construct() {
    d.builder.BuildWalls()
    d.builder.BuildRoof()
    d.builder.BuildWindows()
}

func NewDirector(builder HouseBuilder) *Director {
    return &Director{builder: builder}
}

func main() {
    woodenBuilder := &WoodenHouseBuilder{}
    director := NewDirector(woodenBuilder)
    director.Construct()

    house := woodenBuilder.GetHouse()
    fmt.Printf("House built with: %s, %s, %s\n", house.walls, house.roof, house.windows)
}

