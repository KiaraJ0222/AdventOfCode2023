package main

import (
	"fmt"
	"math"
	"strings"
)

func Day4Part1() {
	lines := ReadFile("day4")
	total := 0

	for line := range lines {
		match := CompareCards(lines, line)
		total = total + CalcuateScore(match)
	}

	fmt.Print("Part 1: ")
	fmt.Println(total)
}

func CalcuateScore(match int) int {
	if match == 0 {
		return 0
	}
	match = match - 1
	float := float64(match)
	power := math.Pow(2, float)
	match = int(power)
	return match
}

func Day4Part2() {
	lines := ReadFile("day4")
	total := 0

	cards := map[int]int{}

	for line := range lines {
		match := CompareCards(lines, line)

		cards[line] = cards[line] + 1
		if match != 0 {
			for i := 1; i <= match; i++ {
				cards[line+i] = cards[line+i] + (1 * (cards[line]))

			}
		}
	}

	for card := range cards {
		total = total + cards[card]
	}

	fmt.Print("Part 2: ")
	fmt.Println(total)
}

func CompareCards(lines []string, line int) int {
	winningLineOfNumbers := (strings.Split((strings.Split(lines[line], ":")[1]), "|")[0])
	elfLineOfNumbers := (strings.Split((strings.Split(lines[line], ":")[1]), "|")[1])
	elfNumbers := strings.Split(elfLineOfNumbers, " ")
	winningNumbers := strings.Split(winningLineOfNumbers, " ")
	match := 0
	for x := range elfNumbers {
		currentElfNumber := strings.TrimSpace(elfNumbers[x])
		if currentElfNumber != "" {
			for y := range winningNumbers {
				currentWinningNumber := strings.TrimSpace(winningNumbers[y])
				if currentWinningNumber != "" {
					if currentWinningNumber == currentElfNumber {
						match = match + 1
					}
				}
			}
		}

	}
	return match
}
