package main

import "fmt"

// Product
type House struct {
    windows string
    doors   string
    roof    string
}

// Builder interface
type HouseBuilder interface {
    SetWindows() HouseBuilder
    SetDoors() HouseBuilder
    SetRoof() HouseBuilder
    Build() House
}

type ConcreteHouseBuilder struct {
	house House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
    return &ConcreteHouseBuilder{}
}

func (b *ConcreteHouseBuilder) SetWindows() HouseBuilder {
    b.house.windows = "wooden windows"
    return b
}

func (b *ConcreteHouseBuilder) SetDoors() HouseBuilder {
    b.house.doors = "wooden doors"
    return b
}

func (b *ConcreteHouseBuilder) SetRoof() HouseBuilder {
    b.house.roof = "tile roof"
    return b
}

func (b *ConcreteHouseBuilder) Build() House {
    return b.house
}

// Director
type Director struct {
    builder HouseBuilder
}

func NewDirector(b HouseBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) Construct() House {
    return d.builder.SetWindows().SetDoors().SetRoof().Build()
}

func (h *House) Show() {
    fmt.Printf("House with %s, %s, and %s roof\n", h.windows, h.doors, h.roof)
}
func main() {
    builder := NewConcreteHouseBuilder()
    director := NewDirector(builder)
    house := director.Construct()
    house.Show()
}

