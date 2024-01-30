package ymachine

import "log"

type Word uint64

var DefaultMemSize Word = 1

// Instructions Opcode constants
const (
	HALT = 0
	NOOP = 1
	INCA = 2
	DECA = 3
)

type ymachine struct {
	P      Word   //'Program Counter' a register on the G-machine which holds the memory address of the next instruction to execute.
	A      Word   // Accumulator
	Memory []Word // Virtual Machine Memory
}

func New() ymachine {
	return ymachine{
		P:      0,
		A:      0,
		Memory: make([]Word, DefaultMemSize),
	}
}

func (g *ymachine) Run() {
	g.Memory = append(g.Memory, HALT)
	log.Println(g.Memory)
	for i := 0; i < len(g.Memory); i++ {
		switch {
		case g.Memory[g.P] == HALT:
			log.Println("HALT")
			g.P++
			return
		case g.Memory[g.P] == NOOP:
			log.Println("NOOP")
		case g.Memory[g.P] == INCA:
			log.Println("INCA")
			g.A++
			log.Println(g.A)
		case g.Memory[g.P] == DECA:
			log.Println("DECA")
			g.A--
			log.Println(g.A)
		}
		g.P++
	}
}

func (g *ymachine) Arithmatic(a, b int) {
	g.Memory = []Word{}
	for i := 0; i < a; i++ {
		g.Memory = append(g.Memory, INCA)
	}
	for i := 0; i < b; i++ {
		g.Memory = append(g.Memory, DECA)
	}
}
