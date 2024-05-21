/* Посетитель — это поведенческий паттерн проектирования, который позволяет добавлять в программу новые операции, 
	не изменяя классы объектов, над которыми эти операции могут выполняться. 

	+:
		- Поведение отделяется от структуры данных. Паттерн позволяет вам добавлять новое поведение к объектам, не изменяя их классы. 
		- Добавление новых операций без изменения классов элементов. Чтобы добавить новую операцию, нужно просто создать новый класс 
		  посетителя, реализующий интерфейс Visitor. 
	-:
		- Доступ к внутреннему состоянию объектов. Посетитель должен иметь доступ к внутренним данным и методам элементов, 
		  что нарушает принцип инкапсуляции.
		- Сильная зависимость. Посетитель и элементы становятся сильно связанными, так как каждый посетитель должен знать 
		  о каждом типе элемента и наоборот.
*/


package main

import "fmt"

// Fruit интерфейс фрукта.
type Fruit interface {
    Accept(visitor Visitor)
}

// Apple структура для яблока.
type Apple struct {
    Color string
}

// Метод Accept, который принимает посетителя и вызывает его метод VisitApple, передавая текущий экземпляр Apple.
func (a *Apple) Accept(visitor Visitor) {
    visitor.VisitApple(a)
}

// Orange структура для апельсина.
type Orange struct {
    Size string
}

// Метод Accept, который принимает посетителя и вызывает его метод VisitOrange, передавая текущий экземпляр Orange.
func (o *Orange) Accept(visitor Visitor) {
    visitor.VisitOrange(o)
}

// Интерфейс, который определяет методы для яблок и апельсинов.
type Visitor interface {
    VisitApple(apple *Apple)
    VisitOrange(orange *Orange)
}

// Cтруктурa FruitInfoVisitor, которая содержит счетчики для яблок и апельсинов.
type FruitInfoVisitor struct {
    AppleCount  int
    OrangeCount int
}

// Метод для структуры FruitInfoVisitor, который увеличивает счетчик яблок.
func (f *FruitInfoVisitor) VisitApple(apple *Apple) {
    f.AppleCount++
}

// Метод для структуры FruitInfoVisitor, который увеличивает счетчик апельсинов.
func (f *FruitInfoVisitor) VisitOrange(orange *Orange) {
    f.OrangeCount++
}

func main() {
    // Создание фруктов.
    fruits := []Fruit{
        &Apple{Color: "Red"},
        &Apple{Color: "Green"},
        &Orange{Size: "Large"},
        &Orange{Size: "Small"},
        &Apple{Color: "Yellow"},
    }

    // Создание посетителя для сбора информации о фруктах.
    fruitInfoVisitor := &FruitInfoVisitor{}

    // В цикле вызываем метод в зависимости от фрукта.
    for _, fruit := range fruits {
        fruit.Accept(fruitInfoVisitor)
    }

    // Вывод информации.
    fmt.Printf("Apples: %d\n", fruitInfoVisitor.AppleCount)
    fmt.Printf("Oranges: %d\n", fruitInfoVisitor.OrangeCount)
}

