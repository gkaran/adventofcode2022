package challenge

import "testing"

const day04TestInput = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

func TestDay04_SolvePartA(t *testing.T) {
	want := "2"
	result := Day04{Input: day04TestInput}.SolvePartA()
	if want != result {
		t.Fatalf("Got %s while expecting %s", result, want)
	}
}

func TestDay04_SolvePartB(t *testing.T) {
	want := "4"
	result := Day04{Input: day04TestInput}.SolvePartB()
	if want != result {
		t.Fatalf("Got %s while expecting %s", result, want)
	}
}
