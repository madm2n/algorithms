package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	for _, testCase := range []struct {
		List   []int
		Item   int
		Result int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			6,
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			10,
			9,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			1,
			0,
		},
	} {
		result := BinarySearch(testCase.List, testCase.Item)

		if result != testCase.Result {
			t.Errorf("list: %v, item: %d, expected: %d, result: %d", testCase.List, testCase.Item, testCase.Result, result)
		}
	}
}
