package rng

import "testing"

func TestIntn(t *testing.T) {
	for i := 0; i < 1000; i++ {
		val := Intn(6)
		if val < 0 || val >= 6 {
			t.Fatalf("Intn(6) returned %d, want [0,6)", val)
		}
	}
}

func TestBool(t *testing.T) {
	trueCount := 0
	for i := 0; i < 1000; i++ {
		if Bool() {
			trueCount++
		}
	}
	if trueCount < 300 || trueCount > 700 {
		t.Fatalf("Bool() returned true %d/1000 times, expected ~500", trueCount)
	}
}

func TestShuffle(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	original := make([]int, len(s))
	copy(original, s)
	Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	sum := 0
	for _, v := range s {
		sum += v
	}
	if sum != 45 {
		t.Fatalf("Shuffle changed elements, sum=%d want 45", sum)
	}
}

func TestIntnPanicsOnZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Intn(0) should panic")
		}
	}()
	Intn(0)
}
