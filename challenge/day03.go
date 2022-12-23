package challenge

import (
	"gkaran/adventofcode2022/utils"
	"strconv"
	"strings"
	"unicode"
)

type Day03 struct {
	Input string
}

func getScore(el rune) int {
	if unicode.IsLower(el) {
		return 1 + int(el-'a')
	} else {
		return 27 + int(el-'A')
	}
}

func (d Day03) SolvePartA() string {
	sum := 0
	for _, rucksack := range strings.Split(d.Input, "\n") {
		first := []rune(rucksack[0 : len(rucksack)/2])
		second := []rune(rucksack[len(rucksack)/2:])
		for _, el := range utils.Intersection(first, second) {
			sum += getScore(el)
		}
	}
	return strconv.Itoa(sum)
}

func (d Day03) SolvePartB() string {
	rucksacks := strings.Split(d.Input, "\n")
	sum := 0
	for i := 0; i < len(rucksacks)-2; i += 3 {
		group := utils.Map(rucksacks[i:i+3], func(i string) []rune {
			return []rune(i)
		})
		for _, el := range utils.Intersection(group...) {
			sum += getScore(el)
		}
	}

	return strconv.Itoa(sum)
}
