/*



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

