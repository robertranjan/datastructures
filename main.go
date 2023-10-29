package main

import (
	"container/heap"
	"fmt"

	myheap "github.com/robertranjan/datastructures/heap"
	linkedlist "github.com/robertranjan/datastructures/linked-list"
	stackandqueue "github.com/robertranjan/datastructures/stack-and-queue"
)

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &myheap.IntHeap{90, 32, 2, 1, 5}
	heap.Init(h)
	fmt.Printf("heap: %v\n", h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

	fmt.Printf("\n\n----------------linkedlist -------------------\n")
	linkedlist.TestQue()
	linkedlist.TestStack()

	fmt.Printf("\n\n----------------stack and queue -------------------\n")
	var S stackandqueue.Stack
	var Q stackandqueue.Queue

	S.Push(250)
	fmt.Printf("S.Pop(): %v", fmt.Sprintln(S.Pop()))

	Q.Enqueue(200)
	fmt.Printf("S.Deque(): %v", fmt.Sprintln(Q.Deque()))

}
