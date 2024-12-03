package main

import (
	"fmt"

	input "github.com/domdierig/adventofcode2024/pkg"
)

func main() {
	lines, err := input.ReadLinesFromFile("./test/day1.txt")

	if err != nil {
		fmt.Printf("Fehler beim Lesen der Datei: %v\n", err)
		return
	}

	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
}
