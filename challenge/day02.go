package challenge

import (
	"strconv"
	"strings"
)

type Day02 struct {
	Input string
}

const RockScore = 1
const PaperScore = 2
const ScissorsScore = 3
const WinScore = 6
const DrawScore = 3
const LoseScore = 0

const Rock = "A"
const Paper = "B"
const Scissors = "C"

func (d Day02) solve(combinations *map[string]int) string {
	var sum = 0
	for _, pair := range strings.Split(d.Input, "\n") {
		sum += (*combinations)[pair]
	}
	return strconv.Itoa(sum)
}

func (d Day02) SolvePartA() string {
	combinations := map[string]int{
		Rock + " X":     DrawScore + RockScore,
		Rock + " Y":     WinScore + PaperScore,
		Rock + " Z":     LoseScore + ScissorsScore,
		Paper + " X":    LoseScore + RockScore,
		Paper + " Y":    DrawScore + PaperScore,
		Paper + " Z":    WinScore + ScissorsScore,
		Scissors + " X": WinScore + RockScore,
		Scissors + " Y": LoseScore + PaperScore,
		Scissors + " Z": DrawScore + ScissorsScore,
	}

	return d.solve(&combinations)
}

func (d Day02) SolvePartB() string {
	combinations := map[string]int{
		Rock + " X":     LoseScore + ScissorsScore,
		Rock + " Y":     DrawScore + RockScore,
		Rock + " Z":     WinScore + PaperScore,
		Paper + " X":    LoseScore + RockScore,
		Paper + " Y":    DrawScore + PaperScore,
		Paper + " Z":    WinScore + ScissorsScore,
		Scissors + " X": LoseScore + PaperScore,
		Scissors + " Y": DrawScore + ScissorsScore,
		Scissors + " Z": WinScore + RockScore,
	}

	return d.solve(&combinations)
}
