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
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v", err)
		return
	}

	fmt.Println("total safe reports: ", totalSafeReports)
}
