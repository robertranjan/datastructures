package slidingwindow

import (
	"fmt"
)

// findShortestWindowMakesN returns result array with 'max of each sub array' with lenth k
// complexity O(n.k)
func FindShortestWindowMakesN(arr []int, k int, n int) []int {
	var result []int
	// var ans int

	// start and end positions
	start, end := 0, k

	// repeat until end is out of array size
	for end <= len(arr) {
		sum := 0
		for _, n := range arr[start:end] {
			sum += n
		}
		// if sum == n {
		// 	ans = 0
		// }
		fmt.Printf("sum of sliding-window %v is %v \n", arr[start:end], sum)
		start++
		end++
		result = append(result, sum)
	}
	return result
}

// findMaxInWindow returns result array with 'max of each sub array' with lenth k
// complexity O(n.k)
func FindMaxInWindow(arr []int, k int) []int {
	var result []int

	// start and end positions
	start, end := 0, k

	// repeat until end is out of array size
	for end <= len(arr) {
		max := arr[start]
		for _, n := range arr[start:end] {
			if n > max {
				max = n
			}
		}
		fmt.Printf("max of sliding-window %v is %v \n", arr[start:end], max)
		start++
		end++
		result = append(result, max)
	}
	return result
}

// findMinInWindow returns result array with 'min of each sub array' with lenth k
// complexity O(n.k)
func FindMinInWindow(arr []int, k int) []int {
	var result []int

	// start and end positions
	start, end := 0, k

	// repeat until end is out of array size
	for end <= len(arr) {
		min := arr[start]
		for _, n := range arr[start:end] {
			if n < min {
				min = n
			}
		}
		fmt.Printf("min of sliding-window %v is %v \n", arr[start:end], min)
		start++
		end++
		result = append(result, min)
	}
	return result
}
