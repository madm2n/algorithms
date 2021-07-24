package breadthfirstsearch

// Queue implementation for search
type Queue [][]string

// Enqueue add new element to queue
func (q *Queue) Enqueue(val []string) {
	*q = append(*q, val)
}

// Dequeue pop element from queue
func (q *Queue) Dequeue() []string {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

// BreadthFirstSearch algorithm implementation
func BreadthFirstSearch(graph map[string][]string, from string, to string) bool {
	queue := Queue{}
	queue.Enqueue(graph[from])

	explored := map[string]bool{}

	for len(queue) > 0 {
		for _, key := range queue.Dequeue() {
			if _, ok := explored[key]; !ok {
				if key == to {
					return true
				} else {
					queue.Enqueue(graph[key])
					explored[key] = true
				}
			}
		}
	}

	return false
}
