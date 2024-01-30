package ymachine

import "log"

type Word uint64

// DefaultMemSize is the number of 64-bit words of memory which will be
// allocated to a new Y-machine by default.
const DefaultMemSize = 1024

// Instructions Opcode constants
const (
	HALT Word = iota + 0
	NOOP
	INCA
	DECA
	SETA
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

func (y *ymachine) Run() {
	for {
		op := y.Fetch()
		switch Word(op) {
		case HALT:
			log.Println("HALT")
			return
		case NOOP:
			log.Println("NOOP")
		case INCA:
			log.Println("INCA")
			y.A++
		case DECA:
			log.Println("DECA")
			y.A--
		case SETA:
			log.Println("SETA")
			y.A = y.Fetch()
		}
		y.Memory = append(y.Memory, HALT)
		log.Println(y.Memory)
	}
}
func (y *ymachine) Fetch() Word {
	op := y.Memory[y.P]
	y.P++
	return op
}

// This implements Addition/Substraction operations where
// a is used for addition and b is used for substraction
// to omit the substraction you can set b = 0 and vice versa
func (y *ymachine) Arithmatic(a, b int) {
	y.Memory = []Word{}
	for i := 0; i < a; i++ {
		y.Memory = append(y.Memory, INCA)
	}
	for i := 0; i < b; i++ {
		y.Memory = append(y.Memory, DECA)
	}
}
