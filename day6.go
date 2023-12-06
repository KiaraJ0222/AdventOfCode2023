package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day6Part1() {
	var times []int
	var distances []int
	var beats []int
	total := 1
	lines := ReadFile("day6")

	unformattedTimes := strings.Split(strings.Split(lines[0], ":")[1], " ")
	unformattedDistances := strings.Split(strings.Split(lines[1], ":")[1], " ")

	for _, time := range unformattedTimes {
		if strings.TrimSpace(time) != "" {
			intTime, _ := strconv.Atoi(time)
			times = append(times, intTime)
		}
	}

	for _, distance := range unformattedDistances {
		if strings.TrimSpace(distance) != "" {
			intDistance, _ := strconv.Atoi(distance)
			distances = append(distances, intDistance)
		}
	}

	for index, time := range times {
		beats = append(beats, 0)
		for i := 0; i <= time; i++ {
			potentialDistance := (time - i) * i
			if potentialDistance > distances[index] {
				beats[index] = beats[index] + 1
			}
		}
	}

	for _, solution := range beats {
		total = total * solution
	}

	fmt.Print("Part 1: ")
	fmt.Println(total)

}

func Day6Part2() {

	total := 0
	lines := ReadFile("day6")

	unformattedTimes := strings.Split(strings.Split(lines[0], ":")[1], " ")
	unformattedDistances := strings.Split(strings.Split(lines[1], ":")[1], " ")

	stringTime := ""
	stringDistance := ""

	for _, time := range unformattedTimes {

		if strings.TrimSpace(time) != "" {
			stringTime = stringTime + time
		}
	}

	for _, distance := range unformattedDistances {

		if strings.TrimSpace(distance) != "" {
			stringDistance = stringDistance + distance
		}
	}

	time, _ := strconv.Atoi(stringTime)
	distance, _ := strconv.Atoi(stringDistance)

	for i := 0; i <= time; i++ {
		potentialDistance := (time - i) * i
		if potentialDistance > distance {
			total++
		}
	}

	fmt.Print("Part 2: ")
	fmt.Println(total)

}
