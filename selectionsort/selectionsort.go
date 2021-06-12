package selectionsort

// SelectionSort algorithm implementation
func SelectionSort(list []int) []int {
	initial, result := make([]int, len(list)), make([]int, 0)

	copy(initial, list)

	for len(initial) > 0 {
		highest := initial[0]
		index := 0

		for i := 1; i < len(initial); i++ {
			guess := initial[i]

			if highest < guess {
				highest = guess
				index = i
			}
		}

		result = append(result, initial[index])
		initial = append(initial[:index], initial[index+1:]...)
	}

	return result
}
