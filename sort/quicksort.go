package main

import "fmt"

// best: O(n log n)
// average: O(n log n)
// worst: O(n^2)
func main() {
	arr := []int{4, 6, 1, 12, 9, 10, 3, 2, 7, 13, 15, 8, 5, 11, 14}
	arr = Quicksort(arr)
	fmt.Println(arr)
}

func Quicksort(arr []int) []int {
	if len(arr) < 2 {return arr}

	var left []int
	var right []int
	pivot := arr[0]

	// partision
	for _, val := range arr[1:] {
		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	arr = append(Quicksort(left), []int{pivot}...)
	arr = append(arr, Quicksort(right)...)
	return arr
}