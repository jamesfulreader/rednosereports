# Day 2 of Advent of Code 2024
## Red Nose Reports
Go implementation for Day 2 of Advent of Code 2024 Historian Histeria [Full Challenge Description](https://adventofcode.com/2024)

## Project Description
My program solves both part 1 and part 2 of the challenge involving reports made up of 5 or more levels:
1. Determine if a report is safe if levels trend positive or negative and if the jump between levels is no greater than 3
2. Add dampener to remove 1 level to determine of that would make a report safe when it would otherwise be unsafe

## Features
- Parse input.txt file containing lines (reports) and levels (numbers)
- Convert strings being read in to integers
- Determine the trending positive or negative direction of a report
- Determine difference between numbers is within safety restrictions not jumping 3 numbers or larger
- Implement dampener function and check sequence function to remove 1 level to see if a report is safe
- Print count of safe reports to the console

## Prerequisites
- `slices` package imported Go 1.21 or higher required
- input.txt file can be downloaded from [Advent of Code 2024](https://adventofcode.com/2024/day/2)

## Installation
```bash
# Clone the repository
git clone https://github.com/jamesfulreader/rednosereports.git
cd rednosereports

# Build the project
go build
```
## input.txt file format
```
7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
```

## Acknowledgments
- [Advent of Code](https://adventofcode.com/) for the original challenge