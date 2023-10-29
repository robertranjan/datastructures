package stackandqueue

import "fmt"

// Stack implement a LIFO
type Stack struct {
	items []int
}

// Queue implement a FIFO
type Queue struct {
	items []int
}

// var S Stack
// var Q Queue

func (s *Stack) Push(i int) {
	fmt.Printf("Pushed %d into stack\n", i)
	s.items = append(s.items, i)
}

func (s *Stack) Pop() (int, bool) {
	item := 0
	if len(s.items) >= 1 {
		item = s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return item, true
	}
	fmt.Printf("Poped %d from stack\n", item)

	return item, false
}

func (q *Queue) Enqueue(i int) {
	fmt.Printf("Enqueing %d into Queue\n", i)
	q.items = append(q.items, i)
}

func (q *Queue) Deque() (int, bool) {
	item := 0
	if len(q.items) >= 1 {
		item = q.items[0]
		q.items = q.items[1:]
		return item, true
	}
	fmt.Printf("Dequeued %d from Queue\n", item)
	return item, false
}
