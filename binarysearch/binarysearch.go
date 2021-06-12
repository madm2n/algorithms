package binarysearch

// BinarySearch algorithm implementation
func BinarySearch(list []int, item int) int {
	low, high := 0, len(list)-1

	for low <= high {
		index := (low + high) / 2
		guess := list[index]

		if guess == item {
			return index
		}

		if guess < item {
			low = index + 1
		} else {
			high = index - 1
		}
	}

	return -1
}
