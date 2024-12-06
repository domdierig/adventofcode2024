package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseLinesToIntArrays(lines []string) ([][]int, error) {
	var result [][]int

	for _, line := range lines {
		parts := strings.Fields(line)

		var numbers []int

		for _, part := range parts {
			num, err := strconv.Atoi(part)

			if err != nil {
				return nil, fmt.Errorf("could not convert '%s' in: '%s'", part, line)
			}

			numbers = append(numbers, int(num))

		}

		result = append(result, numbers)
	}

	return result, nil
}
