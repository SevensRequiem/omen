package cmd

import "testing"

func TestParseDice_Default(t *testing.T) {
	n, sides, mod, err := parseDice("")
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 || sides != 20 || mod != 0 {
		t.Fatalf("got n=%d sides=%d mod=%d, want 1d20+0", n, sides, mod)
	}
}

func TestParseDice_Simple(t *testing.T) {
	n, sides, mod, err := parseDice("2d6")
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 || sides != 6 || mod != 0 {
		t.Fatalf("got n=%d sides=%d mod=%d, want 2d6+0", n, sides, mod)
	}
}

func TestParseDice_NoCount(t *testing.T) {
	n, sides, mod, err := parseDice("d100")
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 || sides != 100 || mod != 0 {
		t.Fatalf("got n=%d sides=%d mod=%d, want 1d100+0", n, sides, mod)
	}
}

func TestParseDice_WithModifier(t *testing.T) {
	n, sides, mod, err := parseDice("4d8+3")
	if err != nil {
		t.Fatal(err)
	}
	if n != 4 || sides != 8 || mod != 3 {
		t.Fatalf("got n=%d sides=%d mod=%d, want 4d8+3", n, sides, mod)
	}
}

func TestParseDice_NegativeModifier(t *testing.T) {
	n, sides, mod, err := parseDice("1d20-5")
	if err != nil {
		t.Fatal(err)
	}
	if n != 1 || sides != 20 || mod != -5 {
		t.Fatalf("got n=%d sides=%d mod=%d, want 1d20-5", n, sides, mod)
	}
}

func TestParseDice_Invalid(t *testing.T) {
	_, _, _, err := parseDice("banana")
	if err == nil {
		t.Fatal("expected error for invalid input")
	}
}
