package quicksort

// QuickSort sorting algorithm implementation
func QuickSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	smaller := []int{}
	bigger := []int{}
	pivot := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] < pivot {
			smaller = append(smaller, list[i])
		} else {
			bigger = append(bigger, list[i])
		}
	}

	return append(append(QuickSort(smaller), pivot), QuickSort(bigger)...)
}
