package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day5(part string) {
	var seed2SoilStart, soil2FertilizerStart, fertilizer2WaterStart, water2LightStart, light2TemperatureStart, temperature2HumidityStart, humidity2LocationStart int

	// seed2LocationMap := make(map[int]int)
	seed2LocationMap := []int{}
	copyOfMap := []int{}
	// copyOfMap := make(map[int]int)

	lines := ReadFile("day5")

	// Create a seeds list
	seedLine := lines[0]
	seeds := strings.Split(strings.Split(seedLine, ":")[1], " ")[1:]
	var intSeeds []int

	if part == "1" {
		for _, seed := range seeds {
			intSeed, _ := strconv.Atoi(seed)
			intSeeds = append(intSeeds, intSeed)
		}
	} else if part == "2" {
		var intSeedsPreRange []int
		for _, seed := range seeds {
			intSeed, _ := strconv.Atoi(seed)
			intSeedsPreRange = append(intSeedsPreRange, intSeed)
		}

		firstPartOfPair := true
		var firstPartOfPairValue int
		for _, seed := range intSeedsPreRange {
			if firstPartOfPair {
				firstPartOfPairValue = seed
				firstPartOfPair = false
			} else {
				for i := 0; i < seed; i++ {
					intSeeds = append(intSeeds, firstPartOfPairValue+i)
				}
				firstPartOfPair = true
			}
		}
	}

	// Find each corresponding index
	for lineIndex, line := range lines {
		if strings.Contains(line, "seed-to-soil map:") {
			seed2SoilStart = lineIndex
		} else if strings.Contains(line, "soil-to-fertilizer map:") {
			soil2FertilizerStart = lineIndex
		} else if strings.Contains(line, "fertilizer-to-water map:") {
			fertilizer2WaterStart = lineIndex
		} else if strings.Contains(line, "water-to-light map:") {
			water2LightStart = lineIndex
		} else if strings.Contains(line, "light-to-temperature map:") {
			light2TemperatureStart = lineIndex
		} else if strings.Contains(line, "temperature-to-humidity map:") {
			temperature2HumidityStart = lineIndex
		} else if strings.Contains(line, "humidity-to-location map:") {
			humidity2LocationStart = lineIndex
		}
	}

	for _, seed := range intSeeds {
		seed2LocationMap = append(seed2LocationMap, seed)
		// seed2LocationMap[seed] = seed
	}

	for k := range seed2LocationMap {
		copyOfMap = append(copyOfMap, seed2LocationMap[k])
	}

	for i := seed2SoilStart + 1; i < soil2FertilizerStart-1; i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}

		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				// delete(seed2LocationMap, seed)
				seed2LocationMap[index] = newLocation
			}
		}
	}

	for k, v := range seed2LocationMap {
		copyOfMap[k] = v
	}

	for i := soil2FertilizerStart + 1; i < fertilizer2WaterStart-1; i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}

		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				// delete(seed2LocationMap, seed)
				seed2LocationMap[index] = newLocation
			}
		}

	}

	for k, v := range seed2LocationMap {
		copyOfMap[k] = v
	}

	for i := fertilizer2WaterStart + 1; i < water2LightStart-1; i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}
		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				seed2LocationMap[index] = newLocation
			}
		}

	}

	for k, v := range seed2LocationMap {
		copyOfMap[k] = v
	}

	for i := water2LightStart + 1; i < light2TemperatureStart-1; i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}

		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				// delete(seed2LocationMap, seed)
				seed2LocationMap[index] = newLocation
			}
		}

	}

	for k, v := range seed2LocationMap {
		copyOfMap[k] = v
	}

	for i := light2TemperatureStart + 1; i < temperature2HumidityStart-1; i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}

		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				// delete(seed2LocationMap, seed)
				seed2LocationMap[index] = newLocation
			}
		}

	}

	for k, v := range seed2LocationMap {
		copyOfMap[k] = v
	}

	for i := temperature2HumidityStart + 1; i < humidity2LocationStart-1; i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}

		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				// delete(seed2LocationMap, seed)
				seed2LocationMap[index] = newLocation
			}
		}

	}

	for k, v := range seed2LocationMap {
		copyOfMap[k] = v
	}

	for i := humidity2LocationStart + 1; i < len(lines); i++ {
		mapping := strings.Split(lines[i], " ")
		var intMapping []int
		for i := range mapping {
			intMap, _ := strconv.Atoi(mapping[i])
			intMapping = append(intMapping, intMap)
		}

		for index := range intSeeds {
			if copyOfMap[index] >= intMapping[1] && copyOfMap[index] <= intMapping[1]+intMapping[2]-1 {
				newLocation := (copyOfMap[index] - intMapping[1]) + intMapping[0]
				seed2LocationMap[index] = newLocation
			}
		}
	}

	answer := -1
	for k := range seed2LocationMap {
		if answer == -1 || answer > seed2LocationMap[k] {
			answer = seed2LocationMap[k]
		}
	}

	fmt.Printf("Answer for %s: ", part)
	fmt.Println(answer)

}
