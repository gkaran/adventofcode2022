package challenge

import (
	"strconv"
	"strings"
)

type Day04 struct {
	Input string
}

func (d Day04) SolvePartA() string {
	return strconv.Itoa(d.countElves(func(e1SMin int, e1SMax int, e2SMin int, e2SMax int) bool {
		return (e1SMin <= e2SMin && e1SMax >= e2SMax) || (e2SMin <= e1SMin && e2SMax >= e1SMax)
	}))
}

func (d Day04) SolvePartB() string {
	return strconv.Itoa(d.countElves(func(e1SMin int, e1SMax int, e2SMin int, e2SMax int) bool {
		return (e1SMin <= e2SMin && e1SMax >= e2SMin) || (e1SMin <= e2SMax && e1SMax >= e2SMax) ||
			(e2SMin <= e1SMin && e2SMax >= e1SMin) || (e2SMin <= e1SMax && e2SMax >= e1SMax)
	}))
}

func (d Day04) countElves(f func(int, int, int, int) bool) (count int) {
	for _, pair := range strings.Split(d.Input, "\n") {
		var elves = strings.Split(pair, ",")
		elf1SectionMin, elf1SectionMax := getElfSections(&elves[0])
		elf2SectionMin, elf2SectionMax := getElfSections(&elves[1])

		if f(elf1SectionMin, elf1SectionMax, elf2SectionMin, elf2SectionMax) {
			count++
		}
	}
	return
}

func getElfSections(sections *string) (sectionMin int, sectionMax int) {
	sectionParts := strings.Split(*sections, "-")
	sectionMin, _ = strconv.Atoi(sectionParts[0])
	sectionMax, _ = strconv.Atoi(sectionParts[1])
	return
}
