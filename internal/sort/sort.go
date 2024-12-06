package sort

import "sort"

func SortInt32Slice(data *[]int) {
	sort.Slice(*data, func(i, j int) bool {
		return (*data)[i] < (*data)[j]
	})
}
