package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2Part2() {
	lines := ReadFile("day2")

	total := 0

	for line := range lines {
		game := strings.Split(lines[line], ":")
		rounds := strings.Split(game[1], ";")

		redMax := 0
		greenMax := 0
		blueMax := 0

		for round := range rounds {
			balls := strings.Split(rounds[round], ",")
			for ball := range balls {
				if strings.Contains(balls[ball], "red") {
					redBalls, _ := strconv.Atoi((strings.Split(balls[ball], " ")[1]))
					if redBalls > redMax {
						redMax = redBalls
					}
				} else if strings.Contains(balls[ball], "green") {
					greenBalls, _ := strconv.Atoi((strings.Split(balls[ball], " ")[1]))
					if greenBalls > greenMax {
						greenMax = greenBalls
					}

				} else if strings.Contains(balls[ball], "blue") {
					blueBalls, _ := strconv.Atoi((strings.Split(balls[ball], " ")[1]))
					if blueBalls > blueMax {
						blueMax = blueBalls
					}
				}
			}
		}

		total = total + (redMax * greenMax * blueMax)
	}

	fmt.Print("Part 2: ")
	fmt.Println(total)
}

func Day2Part1() {
	lines := ReadFile("day2")
	redMax := 12
	greenMax := 13
	blueMax := 14

	total := 0

	for line := range lines {
		game := strings.Split(lines[line], ":")

		gameNumber, _ := strconv.Atoi(strings.Split(game[0], " ")[1])

		rounds := strings.Split(game[1], ";")

		valid := true

		for round := range rounds {
			balls := strings.Split(rounds[round], ",")
			for ball := range balls {
				if strings.Contains(balls[ball], "red") {
					redBalls, _ := strconv.Atoi((strings.Split(balls[ball], " ")[1]))
					if redBalls > redMax {
						valid = false
					}
				} else if strings.Contains(balls[ball], "green") {
					greenBalls, _ := strconv.Atoi((strings.Split(balls[ball], " ")[1]))
					if greenBalls > greenMax {
						valid = false
					}

				} else if strings.Contains(balls[ball], "blue") {
					blueBalls, _ := strconv.Atoi((strings.Split(balls[ball], " ")[1]))
					if blueBalls > blueMax {
						valid = false
					}
				}
			}

		}
		if valid {
			total = total + gameNumber
		}
	}

	fmt.Print("Part 1: ")
	fmt.Println(total)
}
