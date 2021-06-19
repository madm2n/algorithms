package dnc

// Count all elements in the slice
func Count(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return 1 + Count(list[1:])
}
