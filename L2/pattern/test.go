package main

import "fmt"

// Подсистема 1
type CPU struct{}

func (c *CPU) Freeze() {
    fmt.Println("Freezing CPU...")
}

func (c *CPU) Jump(position int) {
    fmt.Printf("Jumping to position %d\n", position)
}

func (c *CPU) Execute() {
    fmt.Println("Executing...")
}

// Подсистема 2
type Memory struct{}

func (m *Memory) Load(position int, data string) {
    fmt.Printf("Loading data '%s' to position %d\n", data, position)
}

// Подсистема 3
type HardDrive struct{}

func (hd *HardDrive) Read(sector int, size int) string {
    fmt.Printf("Reading %d bytes from sector %d\n", size, sector)
    return "some data"
}

// Фасад
type Computer struct {
    cpu     *CPU
    memory  *Memory
    hardDrive *HardDrive
}

func NewComputer() *Computer {
    return &Computer{
        cpu:     &CPU{},
        memory:  &Memory{},
        hardDrive: &HardDrive{},
    }
}

func (c *Computer) Start() {
    c.cpu.Freeze()
    c.memory.Load(0, c.hardDrive.Read(0, 1024))
    c.cpu.Jump(0)
    c.cpu.Execute()
}

func main() {
    computer := NewComputer()
    computer.Start()
}

