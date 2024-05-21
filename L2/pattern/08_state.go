/*
	Паттерн Состояние — это поведенческий паттерн проектирования,
	который позволяет объекту изменять свое поведение при изменении его внутреннего состояния.

	+:
		- Легко добавлять новые состояния и связанные с ними поведения.
		- Логика, зависящая от состояния, инкапсулируется в отдельных классах состояний, что упрощает понимание и изменение этой логики.
	
	-:
		- Если переходы между состояниями сложны, их управление может стать трудоемким и запутанным.
*/

package main

import "fmt"

// Интерфейс для различных состояний плеера.
type State interface {
	Play(p *Player)
	Stop(p *Player)
}

// Представляет контекст, который хранит текущее состояние.
type Player struct {
	state State
}

// Cоздаёт новый плеер в состоянии остановлен.
func NewPlayer() *Player {
	return &Player{state: &Stopped{}}
}

// Yстанавливает новое состояние для плеера.
func (p *Player) SetState(state State) {
	p.state = state
}

// Bызывает метод Play текущего состояния.
func (p *Player) Play() {
	p.state.Play(p)
}

// Bызывает метод Stop текущего состояния.
func (p *Player) Stop() {
	p.state.Stop(p)
}

// Cостояние остановки.
type Stopped struct{}

func (s *Stopped) Play(p *Player) {
	fmt.Println("Starting playback.")
	p.SetState(&Playing{})
}

func (s *Stopped) Stop(p *Player) {
	fmt.Println("Player is already stopped.")
}

// Cостояние воспроизведения.
type Playing struct{}

func (pl *Playing) Play(p *Player) {
	fmt.Println("Already playing.")
}

func (pl *Playing) Stop(p *Player) {
	fmt.Println("Stopping playback.")
	p.SetState(&Stopped{})
}

func main() {
	player := NewPlayer()

	player.Play()
	player.Play()
	player.Stop()
	player.Stop()
}

