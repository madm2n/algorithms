package dnc

import "testing"

func TestMax(t *testing.T) {
	for _, testCase := range []struct {
		List   []int
		Result int
	}{
		{
			[]int{},
			0,
		},
		{
			[]int{1, 2, 4, 5, 10, 22},
			22,
		},
		{
			[]int{-20, -100, -200, 1, 2, 4, 5, 6},
			6,
		},
		{
			[]int{-20, -100, -200, -1, -2, -3, -4},
			-1,
		},
		{
			[]int{-20, -100, 0, -200, -1, -2, -3, -4},
			0,
		},
	} {
		result := Max(testCase.List)

		if result != testCase.Result {
			t.Errorf("list: %v, expected: %d, result: %d", testCase.List, testCase.Result, result)
		}
	}
}
