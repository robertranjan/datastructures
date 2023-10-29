package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	next *Node
}

func (s *Stack) Push(i int) {
	newNode := &Node{value: i}
	newNode.next = s.next
	s.next = newNode
	fmt.Printf("pushed %v into stack\n", i)
}

func (s *Stack) Pop() (int, bool) {
	if s.next == nil {
		return 0, false
	}
	currentNode := s.next
	item := currentNode.value
	s.next = currentNode.next
	return item, true
}

func TestStack() {
	fmt.Printf("---------------- Stack LIFO -----------------\n")
	s := &Stack{next: nil}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	for {
		item, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Printf("poped %v\n", item)
	}
}

type Queue struct {
	front *Node
	rear  *Node
}

func (q *Queue) Enque(i int) {
	newNode := &Node{value: i}
	fmt.Printf("enquing %v into queue\n", i)
	if q.front == nil {
		q.front = newNode
		q.rear = newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

func (q *Queue) Deque() (int, bool) {
	var item int
	if q.front == nil {
		return item, false
	}
	item = q.front.value
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}
	return item, true
}
func TestQue() {
	fmt.Printf("---------------- Queue FIFO -----------------\n")
	q := &Queue{front: nil, rear: nil}
	q.Enque(10)
	q.Enque(20)
	q.Enque(30)
	q.Enque(32)

	for {
		item, ok := q.Deque()
		if !ok {
			break
		}

		fmt.Printf("q.Deque(): %v\n", item)
	}
}
