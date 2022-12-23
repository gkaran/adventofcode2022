package challenge

import "testing"

const day02TestInput = `A Y
B X
C Z`

func TestDay02_SolvePartA(t *testing.T) {
	want := "15"
	result := Day02{Input: day02TestInput}.SolvePartA()
	if want != result {
		t.Fatalf("Got %s while expecting %s", result, want)
	}
}

func TestDay02_SolvePartB(t *testing.T) {
	want := "12"
	result := Day02{Input: day02TestInput}.SolvePartB()
	if want != result {
		t.Fatalf("Got %s while expecting %s", result, want)
	}
}
