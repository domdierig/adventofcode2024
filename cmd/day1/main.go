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

	var leftList []int
	var rightList []int

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

	differenceresult := 0

	for i := 0; i < len(leftList); i++ {
		difference := int(math.Abs(float64(leftList[i] - rightList[i])))
		differenceresult += difference
	}

	fmt.Printf("the distrance result is: %d\n", differenceresult)

	similarityresult := 0

	for _, leftnumber := range leftList {
		count := 0
		for _, rightnumber := range rightList {
			if leftnumber == rightnumber {
				count++
			}
		}

		similarityresult += (leftnumber * count)
	}

	fmt.Printf("the similarity result is: %d\n", similarityresult)
}
