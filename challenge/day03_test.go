package challenge

import "testing"

const day03TestInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestDay03_SolvePartA(t *testing.T) {
	want := "157"
	result := Day03{Input: day03TestInput}.SolvePartA()
	if want != result {
		t.Fatalf("Got %s while expecting %s", result, want)
	}
}

func TestDay03_SolvePartB(t *testing.T) {
	want := "70"
	result := Day03{Input: day03TestInput}.SolvePartB()
	if want != result {
		t.Fatalf("Got %s while expecting %s", result, want)
	}
}
