package challenge

import "testing"

const day01TestInput = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestDay01_SolvePartA(t *testing.T) {
	if "24000" != (Day01{Input: day01TestInput}.SolvePartA()) {
		t.Fail()
	}
}

func TestDay01_SolvePartB(t *testing.T) {
	if "45000" != (Day01{Input: day01TestInput}.SolvePartB()) {
		t.Fail()
	}
}
