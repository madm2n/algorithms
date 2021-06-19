package dnc

// Max find max number in slice
func Max(list []int) int {
	if len(list) == 0 {
		return 0
	}

	if len(list) == 1 {
		return list[0]
	}

	max := Max(list[1:])

	if list[0] > max {
		return list[0]
	}

	return max
}
