package calculator

import "testing"

func TestAdditon(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("Expected 1 (+) 2 to equal 3")
	}
	if Additon(-1, -2) != -3 {
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubstract(t *testing.T) {
	result := Substract(1, 2)
	if result != -1 {
		t.Errorf("Expected 1 - 2 to equal -1, but got %d", result)
	}

	result = Substract(-1, -2)
	if result != 1 {
		t.Errorf("Expected -1 - (-2) to equal 1, but got %d", result)
	}
}

func TestMult(t *testing.T) {
	if Mult(1, 2) != 2 {
		t.Error("Expected 1 (*) 2 to equal 2")
	}
	if Mult(-1, -2) != 2 {
		t.Error("Expected -1 (*) -2 to equal 2")
	}
}

func TestDiv(t *testing.T) {
	result := Div(2, 1)
	if result != 2 {
		t.Errorf("Expected 2 / 1 to equal 2, but got %d", result)
	}

	result = Div(-2, -1)
	if result != 2 {
		t.Errorf("Expected -2 / (-1) to equal 2, but got %d", result)
	}
}
