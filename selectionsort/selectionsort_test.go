package selectionsort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	for _, testCase := range []struct {
		List   []int
		Result []int
	}{
		{
			[]int{10, 9, 2, 1, 5, 8},
			[]int{10, 9, 8, 5, 2, 1},
		},

		{
			[]int{5, 8, 9, 10, 2, 1},
			[]int{10, 9, 8, 5, 2, 1},
		},
	} {
		result := SelectionSort(testCase.List)

		if len(result) != len(testCase.Result) {
			t.Fatalf("list: %v, expected: %v, result: %v", testCase.List, testCase.Result, result)
		}

		for i := 0; i < len(testCase.Result); i++ {
			if testCase.Result[i] != result[i] {
				t.Fatalf("list: %v, expected: %v, result: %v", testCase.List, testCase.Result, result)
			}
		}
	}
}
