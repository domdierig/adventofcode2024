package parse

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseLinesToIntArrays(lines []string) ([][]int32, error) {
	var result [][]int32

	for _, line := range lines {
		parts := strings.Fields(line)

		var numbers []int32

		for _, part := range parts {
			num, err := strconv.Atoi(part)

			if err != nil {
				return nil, fmt.Errorf("could not convert '%s' in: '%s'", part, line)
			}

			numbers = append(numbers, int32(num))

		}

		result = append(result, numbers)
	}

	return result, nil
}
