package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, testCase := range []struct {
		List   []int
		Result []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{5, 3, 4, 1, 2},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{-1, -2, -3, -4, -5},
			[]int{-5, -4, -3, -2, -1},
		},
		{
			[]int{-5, -3, -2, -4, -1},
			[]int{-5, -4, -3, -2, -1},
		},
	} {
		result := QuickSort(testCase.List)

		if !reflect.DeepEqual(result, testCase.Result) {
			t.Errorf("list: %v, expected: %d, result: %d", testCase.List, testCase.Result, result)
		}
	}
}
