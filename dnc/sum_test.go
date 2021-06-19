package dnc

import "testing"

func TestSum(t *testing.T) {
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
			1 + 2 + 4 + 5 + 10 + 22,
		},
		{
			[]int{-20, -100, -200, 1, 2, 4, 5, 6},
			-20 - 100 - 200 + 1 + 2 + 4 + 5 + 6,
		},
		{
			[]int{-20, -100, -200, -1, -2, -3, -4},
			-20 - 100 - 200 - 1 - 2 - 3 - 4,
		},
		{
			[]int{-20, -100, 0, -200, -1, -2, -3, -4},
			-20 + -100 + 0 - 200 - 1 - 2 - 3 - 4,
		},
	} {
		result := Sum(testCase.List)

		if result != testCase.Result {
			t.Errorf("list: %v, expected: %d, result: %d", testCase.List, testCase.Result, result)
		}
	}
}
