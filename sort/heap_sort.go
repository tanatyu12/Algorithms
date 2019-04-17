package main

import "fmt"

// best: O(n log n)
// average: O(n log n)
// worst: O(n log n)
func main() {
	arr := []int{4, 6, 1, 12, 9, 10, 3, 2, 7, 13, 15, 8, 5, 11, 14}
	BuildHeap(&arr);

	for i := len(arr)-1; i > 0; i-- {
		tmp := arr[i]
		arr[i] = arr[0]
		arr[0] = tmp
		Heapify(&arr, 0, i)
	}
	fmt.Println(arr)
}

func BuildHeap(ap *[]int) {
	n := len(*ap)
	for i := n/2; i >= 0; i-- {
		Heapify(ap, i, n)
	}
}

func Heapify(ap *[]int, idx int, max int) {
	largest := idx
	left := idx * 2 + 1
	right := idx * 2 + 2

	if left < max && (*ap)[left] > (*ap)[largest] {
		largest = left
	}
	if right < max && (*ap)[right] > (*ap)[largest] {
		largest = right
	}
	if idx != largest {
		tmp := (*ap)[idx]
		(*ap)[idx] = (*ap)[largest]
		(*ap)[largest] = tmp
		Heapify(ap, largest, max)
	}
} 

