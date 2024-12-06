package main

import (
	"fmt"

	"github.com/domdierig/adventofcode2024/internal/input"
	"github.com/domdierig/adventofcode2024/internal/parse"
)

func main() {
	lines, err := input.ReadLinesFromFile("./test/day2.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	lineNumbers, err := parse.ParseLinesToIntArrays(lines)

	if err != nil {
		fmt.Print(err)
		return
	}

	safeLines := 0

	strategies := map[string]func(a int, b int) bool{
		"inc": shouldIncreaseButNotMoreThan3,
		"dec": shouldDecreaseButNotMoreThan3,
	}

	for _, numbers := range lineNumbers {
		isSafe := true

		chosenStrategy := strategies["dec"]

		if numbers[0] < numbers[1] {
			chosenStrategy = strategies["inc"]
		}

		if numbers[0] != numbers[1] {
			for i := 1; i < len(numbers); i++ {
				if chosenStrategy(numbers[i], numbers[i-1]) {

				} else {
					isSafe = false
				}
			}
		} else {
			isSafe = false
		}

		if isSafe {
			safeLines++
		}
	}

	fmt.Printf("count of safe lines is: %d\n", safeLines)
}

func shouldIncreaseButNotMoreThan3(current int, predecessor int) bool {
	if current > predecessor && current-predecessor <= 3 {
		return true
	}

	return false
}

func shouldDecreaseButNotMoreThan3(current int, predecessor int) bool {
	return shouldIncreaseButNotMoreThan3(predecessor, current)
}
