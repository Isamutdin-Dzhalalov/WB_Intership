*/
	Паттерн "Команда" (Command) — это поведенческий паттерн проектирования, который превращает запросы в объекты, 
	позволяя передавать их как параметры.

	+:
	  -	Инкапсуляция запросов: Запросы превращаются в объекты, что позволяет легко передавать, хранить и управлять ими.
	  - Легкость расширения: Новые команды можно добавлять без изменения существующего кода.

	-:
	  - Усложнение кода: Увеличивает количество классов и усложняет структуру приложения.
	  - Избыточность: Может добавлять избыточные накладные расходы, особенно для простых операций.

	Примеры:
		Паттерн "Команда" используется для создания макросов — последовательностей команд, 
		которые можно записать и воспроизвести. Это особенно полезно в приложениях для автоматизации повторяющихся задач.
*/

package main

import "fmt"

// Command интерфейс команды.
type Command interface {
    Execute()
}

// PrintHello команда для печати hello сообщения.
type PrintHello struct{}

// Реализует метод для команды PrintHello, который печает Hello.
func (c *PrintHello) Execute() {
    fmt.Println("Hello!")
}

// PrintGoodbye команда для печати goodbye сообщения.
type PrintGoodbye struct{}

// Реализует метод для команды PrintGoodbye, который печает Goodbye.
func (c *PrintGoodbye) Execute() {
    fmt.Println("Goodbye!")
}

// Invoker структура, вызывающая команды.
type Invoker struct {
    command Command
}

// SetCommand устанавливает команду для выполнения.
func (i *Invoker) SetCommand(command Command) {
    i.command = command
}

// Run выполняет установленную команду.
func (i *Invoker) Run() {
    i.command.Execute()
}

func main() {
    // Создание команд.
    hello := &PrintHello{}
    goodbye := &PrintGoodbye{}

    // Создание инициатора.
    invoker := &Invoker{}

    // Выполнение команды печати hello.
    invoker.SetCommand(hello)
    invoker.Run()

    // Выполнение команды печати goodbye.
    invoker.SetCommand(goodbye)
    invoker.Run()
}

