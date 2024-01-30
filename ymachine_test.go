package ymachine_test

import (
	"testing"

	"github.com/MhmoudGit/ymachine"
)

func TestNew(t *testing.T) {
	t.Parallel()
	g := ymachine.New()
	var wantP, wantA ymachine.Word = 0, 0
	if wantP != g.P {
		t.Errorf("want initial P value %d, got %d", wantP, g.P)
	}
	if wantA != g.A {
		t.Errorf("want initial P value %d, got %d", wantA, g.A)
	}
	var wantMemValue ymachine.Word = 0
	gotMemValue := g.Memory[ymachine.DefaultMemSize-1]
	if wantMemValue != gotMemValue {
		t.Errorf("want last memory location to contain %d, got %d", wantMemValue, gotMemValue)
	}
}

func TestHalt(t *testing.T) {
	t.Parallel()
	g := ymachine.New()
	g.Run()
	var wantP ymachine.Word = 1
	if wantP != g.P {
		t.Errorf("want P == %d, got, got %d", wantP, g.P)
	}
}

func TestNoop(t *testing.T) {
	t.Parallel()
	g := ymachine.New()
	g.Memory[0] = ymachine.NOOP
	g.Run()
	var wantP ymachine.Word = 2
	if wantP != g.P {
		t.Errorf("want P == %d, got, got %d", wantP, g.P)
	}
}

func TestINCA(t *testing.T) {
	t.Parallel()
	g := ymachine.New()
	g.Memory[0] = ymachine.INCA
	g.Run()
	var wantA ymachine.Word = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got, got %d", wantA, g.A)
	}
}

func TestDECA(t *testing.T) {
	t.Parallel()
	g := ymachine.New()
	g.A = 2
	g.Memory[0] = ymachine.DECA
	g.Run()
	var wantA ymachine.Word = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got, got %d", wantA, g.A)
	}
}

func TestSubtracting(t *testing.T) {
	t.Parallel()
	g := ymachine.New()
	g.Arithmatic(3, 2)
	g.Run()
	var wantA ymachine.Word = 1
	if wantA != g.A {
		t.Errorf("want A == %d, got, got %d", wantA, g.A)
	}
}
