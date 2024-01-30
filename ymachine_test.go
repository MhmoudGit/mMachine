package ymachine_test

import (
	"testing"

	"github.com/MhmoudGit/ymachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	var wantP, wantA ymachine.Word = 0, 0
	if wantP != y.P {
		t.Errorf("want initial P value %d, got %d", wantP, y.P)
	}
	if wantA != y.A {
		t.Errorf("want initial P value %d, got %d", wantA, y.A)
	}
	var wantMemValue ymachine.Word = 0
	gotMemValue := y.Memory[ymachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHalt(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	y.Run()
	var wantP ymachine.Word = 1
	if wantP != y.P {
		t.Errorf("want P == %d, got, got %d", wantP, y.P)
	}
}

func TestNoop(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	y.Memory[0] = ymachine.NOOP
	y.Run()
	var wantP ymachine.Word = 2
	if wantP != y.P {
		t.Errorf("want P == %d, got, got %d", wantP, y.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	y.Memory[0] = ymachine.INCA
	y.Run()
	var wantA ymachine.Word = 1
	if wantA != y.A {
		t.Errorf("want A == %d, got, got %d", wantA, y.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	y.A = 2
	y.Memory[0] = ymachine.DECA
	y.Run()
	var wantA ymachine.Word = 1
	if wantA != y.A {
		t.Errorf("want A == %d, got, got %d", wantA, y.A)
	}
}

func TestArithmatic(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	y.Arithmatic(3, 2)
	y.Run()
	var wantA ymachine.Word = 1
	if wantA != y.A {
		t.Errorf("want A == %d, got, got %d", wantA, y.A)
	}
}

func TestSETA(t *testing.T) {
	t.Parallel()
	y := ymachine.New()
	y.Memory[0] = ymachine.SETA
	y.Memory = append(y.Memory, 5)
	y.Run()
	var wantA ymachine.Word = 5
	if wantA != y.A {
		t.Errorf("want A == %d, got, got %d", wantA, y.A)
	}
	var wantP ymachine.Word = 3
	if wantP != y.P {
		t.Errorf("want P == %d, got, got %d", wantP, y.P)
	}
}
