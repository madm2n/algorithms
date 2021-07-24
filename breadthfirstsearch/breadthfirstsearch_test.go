package breadthfirstsearch

import (
	"testing"
)

func TestBreadthFirstSearch(t *testing.T) {
	for _, testCase := range []struct {
		Graph  map[string][]string
		From   string
		To     string
		Result bool
	}{
		{
			map[string][]string{
				"you":     {"alice", "bob", "claire"},
				"bob":     {"anuj", "peggy", "stephan"},
				"alice":   {"thom", "jonny"},
				"claire":  {},
				"anuj":    {},
				"peggy":   {},
				"thom":    {},
				"jonny":   {},
				"stephan": {"jim"},
				"jim":     {},
			},
			"you",
			"jim",
			true,
		},
		{
			map[string][]string{
				"you":     {"alice", "bob", "claire"},
				"bob":     {"anuj", "peggy", "stephan"},
				"alice":   {"thom", "jonny"},
				"claire":  {},
				"anuj":    {},
				"peggy":   {},
				"thom":    {},
				"jonny":   {},
				"stephan": {"jim"},
				"jim":     {},
			},
			"jim",
			"you",
			false,
		},
	} {
		result := BreadthFirstSearch(testCase.Graph, testCase.From, testCase.To)

		if result != testCase.Result {
			t.Errorf("graph: %v, from: %s, to: %s, expected: %v, result: %v", testCase.Graph, testCase.From, testCase.To, testCase.Result, result)
		}
	}
}
