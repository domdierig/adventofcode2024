package main

import (
	"fmt"
	"math"

	"github.com/domdierig/adventofcode2024/internal/input"
	"github.com/domdierig/adventofcode2024/internal/parse"
	"github.com/domdierig/adventofcode2024/internal/sort"
)

func main() {
	lines, err := input.ReadLinesFromFile("./test/day1.txt")

	if err != nil {
		fmt.Print(err)
		return
	}

	lineNumbers, err := parse.ParseLinesToIntArrays(lines)

	if err != nil {
		fmt.Print(err)
		return
	}

	var leftList []int32
	var rightList []int32

	for _, numbers := range lineNumbers {

		leftList = append(leftList, numbers[0])
		rightList = append(rightList, numbers[1])
	}

	sort.SortInt32Slice(&leftList)
	sort.SortInt32Slice(&rightList)

	if len(leftList) != len(rightList) {
		fmt.Print("lists have not same size")
		return
	}

	var result int32 = 0

	for i := 0; i < len(leftList); i++ {
		difference := int32(math.Abs(float64(leftList[i] - rightList[i])))
		result += difference
	}

	fmt.Printf("the result is: %d\n", result)
}
