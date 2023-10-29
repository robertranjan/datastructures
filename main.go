package main

import (
	"container/heap"
	"fmt"

	myheap "github.com/robertranjan/datastructures/heap"
	"github.com/robertranjan/datastructures/linkedlist"
	"github.com/robertranjan/datastructures/slidingwindow"
	"github.com/robertranjan/datastructures/stackandqueue"
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

	fmt.Printf("\n\n----------------sliding window-------------------\n")
	arr := []int{1, 2, 3, 4, 5, 6, 3, 2, 1, 4, 2, 1, 6, 32, 3}
	fmt.Printf("\nMax of sliding window\n")
	fmt.Println("result: ", slidingwindow.FindMaxInWindow(arr, 3))
	fmt.Printf("\nMin of sliding window\n")
	fmt.Println("result: ", slidingwindow.FindMinInWindow(arr, 3))
	fmt.Printf("\nfindShortestWindowMakesN of sliding window\n")
	fmt.Println("result: ", slidingwindow.FindShortestWindowMakesN(arr, 3, 5))
}
