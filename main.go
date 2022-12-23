package main

import (
	"errors"
	"flag"
	"fmt"
	"gkaran/adventofcode2022/challenge"
	"gkaran/adventofcode2022/challenge/common"
	"log"
	"os"
)

func main() {
	day := flag.Int("day", 0, "challenge day")
	part := flag.String("part", "", "day part (a or b)")
	flag.Parse()

	input := getDayInput(*day)
	d, _ := getDayChallenge(day, &input)

	printDayResult(*day, *part, solveDayPart(&d, part))
}

func getDayChallenge(day *int, input *string) (common.Day, error) {
	switch *day {
	case 1:
		return challenge.Day01{Input: *input}, nil
	case 2:
		return challenge.Day02{Input: *input}, nil
	case 3:
		return challenge.Day03{Input: *input}, nil
	case 4:
		return challenge.Day04{Input: *input}, nil
	default:
		return nil, errors.New("NO IMPLEMENTATION FOR DAY EXISTS")
	}
}

func solveDayPart(day *common.Day, part *string) string {
	if *part == "a" {
		return (*day).SolvePartA()
	}
	return (*day).SolvePartB()
}

func printDayResult(day int, part string, result string) {
	fmt.Printf("Day = %d and part = %s: %s\n", day, part, result)
}

func getDayInput(day int) string {
	content, err := os.ReadFile(fmt.Sprintf("./inputs/day%02d.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
