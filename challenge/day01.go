package challenge

import (
	"gkaran/adventofcode2022/utils"
	"sort"
	"strconv"
	"strings"
)

type Day01 struct {
	Input string
}

func (d Day01) SolvePartA() string {
	var m = 0
	for _, elf := range strings.Split(d.Input, "\n\n") {
		var calories = 0
		for _, c := range strings.Split(elf, "\n") {
			cal, _ := strconv.Atoi(c)
			calories += cal
		}
		m = utils.Max(m, calories)
	}
	return strconv.Itoa(m)
}

func (d Day01) SolvePartB() string {
	var elfCalories []int
	for _, elf := range strings.Split(d.Input, "\n\n") {
		var calories = 0
		for _, c := range strings.Split(elf, "\n") {
			cal, _ := strconv.Atoi(c)
			calories += cal
		}
		elfCalories = append(elfCalories, calories)
	}

	var total = 0
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))
	for _, cal := range elfCalories[0:3] {
		total += cal
	}

	return strconv.Itoa(total)
}
