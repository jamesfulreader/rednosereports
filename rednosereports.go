package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	minDiff = 1
	maxDiff = 3
)

type direction int

const (
	positive direction = iota
	negative
)

func checkSequence(report []string) bool {
	if len(report) < 2 {
		return false
	}

	levels := make([]int, len(report))
	for i, level := range report {
		num, err := strconv.Atoi(level)
		if err != nil {
			return false
		}
		levels[i] = num
	}

	// Check if strictly increasing
	isValid := true
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff <= 0 || diff > 3 {
			isValid = false
			break
		}
	}
	if isValid {
		return true
	}

	// Check if strictly decreasing
	isValid = true
	for i := 1; i < len(levels); i++ {
		diff := levels[i-1] - levels[i]
		if diff <= 0 || diff > 3 {
			isValid = false
			break
		}
	}
	return isValid
}

func dampener(reports []string) bool {

	if checkSequence(reports) {
		return true
	}

	// Try removing each number (including first and last)
	for i := 0; i < len(reports); i++ {
		newReports := make([]string, 0, len(reports)-1)
		// get all levels up to level being removed
		newReports = append(newReports, reports[:i]...)
		// get all levels after the level being removed
		newReports = append(newReports, reports[i+1:]...)

		if checkSequence(newReports) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println("Red nose reports!")
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer input.Close()

	totalSafeReports := 0
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		reports := strings.Fields(scanner.Text())
		if len(reports) < 2 {
			fmt.Println("Warning: skipping invalid report line")
			continue
		}

		levelOne, err := strconv.Atoi(reports[0])
		if err != nil {
			fmt.Printf("Warning: %v\n", err)
			continue
		}

		levelTwo, err := strconv.Atoi(reports[1])
		if err != nil {
			fmt.Printf("Warning: %v\n", err)
			continue
		}

		initialDirection := positive
		if levelOne > levelTwo {
			initialDirection = negative
		}

		previousLevel := levelOne
		isReportSafe := true

		for _, level := range reports[1:] {
			currentLevel, err := strconv.Atoi(level)
			if err != nil {
				fmt.Printf("Warning: %v\n", err)
				isReportSafe = false
				break
			}

			difference := math.Abs(float64(currentLevel) - float64(previousLevel))

			if difference < minDiff || difference > maxDiff {
				isReportSafe = false
				break
			}

			if initialDirection == positive {
				if currentLevel <= previousLevel {
					isReportSafe = false
					break
				}
			} else {
				if currentLevel >= previousLevel {
					isReportSafe = false
					break
				}
			}

			previousLevel = currentLevel
		}
		if isReportSafe {
			totalSafeReports++
		} else if dampener(reports) {
			totalSafeReports++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v", err)
		return
	}

	fmt.Println("total safe reports: ", totalSafeReports)
}
