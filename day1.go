package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Day1Part1() {
	total := 0
	lines := ReadFile("day1")

	for x := range lines {
		var numbers []int
		chars := strings.Split(lines[x], "")

		for y := range chars {
			if num, err := strconv.Atoi(chars[y]); err == nil {
				numbers = append(numbers, num)
			}
		}

		thisLineTotal := strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[len(numbers)-1])
		thisLineTotalNum, _ := strconv.Atoi(thisLineTotal)
		total = total + thisLineTotalNum
	}

	fmt.Print("Part 1: ")
	fmt.Println(total)
}

func Day1Part2() {

	total := 0
	lines := ReadFile("day1")

	numberPattern := "one|two|three|four|five|six|seven|eight|nine|0|1|2|3|4|5|6|7|8|9"
	numberPatternReversed := ReverseString(numberPattern)

	regexForward, _ := regexp.Compile(numberPattern)
	regexBackward, _ := regexp.Compile(numberPatternReversed)

	for x := range lines {
		// var numbers []string
		forwardCheck := []byte(lines[x])
		first := regexForward.Find(forwardCheck)

		backwardsCheck := []byte(ReverseString(lines[x]))
		last := regexBackward.Find(backwardsCheck)

		thisLineTotal := strconv.Itoa(ConvertStringNumToInt(string(first))) + strconv.Itoa(ConvertStringNumToInt(string(last)))
		thisLineTotalNum, _ := strconv.Atoi(thisLineTotal)
		total = total + thisLineTotalNum
	}

	fmt.Print("Part 2: ")
	fmt.Println(total)
}

func ReverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func ConvertStringNumToInt(found string) int {
	if num, err := strconv.Atoi(found); err == nil {
		return num
	} else {
		switch found {
		case "one":
			return 1
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		case "eno":
			return 1
		case "owt":
			return 2
		case "eerht":
			return 3
		case "ruof":
			return 4
		case "evif":
			return 5
		case "xis":
			return 6
		case "neves":
			return 7
		case "thgie":
			return 8
		case "enin":
			return 9
		}
	}

	return -1
}
