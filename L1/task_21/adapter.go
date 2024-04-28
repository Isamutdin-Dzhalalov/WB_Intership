package main

import "fmt"

// Интерфейс, который мы хотим адаптировать.
type OldInterface interface {
    OldMethod() string
}

// Класс, реализующий старый интерфейс.
type OldClass struct{}

func (oc *OldClass) OldMethod() string {
    return "Old Method"
}

// Новый интерфейс, к которому мы хотим адаптировать.
type NewInterface interface {
    NewMethod() string
}

// Адаптер, реализующий новый интерфейс и использующий старый.
type Adapter struct {
    old OldInterface
}

func (a *Adapter) NewMethod() string {
    return a.old.OldMethod()
}

func main() {
    oldClass := &OldClass{}
    adapter := &Adapter{old: oldClass}

    // Использование адаптера для вызова метода нового интерфейса.
    fmt.Println(adapter.NewMethod())
}

