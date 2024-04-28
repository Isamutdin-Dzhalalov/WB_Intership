package main

import (
	"fmt"
)

// Структура с двумя полями
type Human struct {
	Name string
	Age int
}

/* Структура с анонимным полем Human, с ипользованием
	типа структуры, позволяет нам обращаться с полям
	структуры "Human" без указания имени, или типа, 
	как показано в функции "printStruct"*/

type Action struct {
	Human
}

/* Использование указателя в качестве параметра получателя,
   что позволяет методу изменять исходные данные, так как он
   получил адрес, а не копию. */

func (a *Action) ini() {
	a.Name = "Isma"
	a.Age = 30
}

/* Передача по значение, данный метод не меняет исходные данные 
	переданного аргумента, так как создаётся копия аргумента. */

func (a Action) printStruct() {
	fmt.Printf("Имя: %s\n", a.Name) // Обращаемся к полю структуры "Human" без указания типа.
	fmt.Printf("Возраст: %d\n", a.Human.Age) // Обращаемся к полю "Human" через указание типа.
}
		

func main() {
	var action Action

	// Вызов метода: "получатель метода"."имя метода()"
	
	action.printStruct()
	action.ini()
	fmt.Println()
	action.printStruct()
}
