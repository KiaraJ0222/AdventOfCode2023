package main

import (
	"os"
	"strings"
)

func ReadFile(file string) []string {
	// Read the file
	file = "./inputs/" + file + ".txt"
	data, _ := os.ReadFile(file)
	// Split text
	input := string(data)
	return (strings.Split(input, "\n"))
}
