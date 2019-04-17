package main

import "fmt"

// best: O(n log n)
// average: O(n log n)
// worst: O(n^2)

// pivot is median-of-three
// apply insertion sort to small partial array
func main() {
	arr := []int{4, 6, 1, 12, 9, 10, 3, 2, 7, 13, 15, 8, 5, 11, 14}

	arr = TransQuicksort(arr)
	fmt.Println(arr)
}

func TransQuicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	var left []int
	var right []int
	pivotIdx := SelectPivot(arr)

	// partision
	for _, val := range arr[:pivotIdx] {
		if val <= arr[pivotIdx] {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	for _, val := range arr[pivotIdx+1:] {
		if val <= arr[pivotIdx] {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	// if array size become less than minSize, switch quicksort to insertion sort
	minSize := 5
	if len(left) < minSize {
		arr = append(InsertionSort(left), []int{arr[pivotIdx]}...)
	} else {
		arr = append(TransQuicksort(left), []int{arr[pivotIdx]}...)
	}
	if len(right) < minSize {
		arr = append(arr, InsertionSort(right)...)
	} else {
		arr = append(arr, TransQuicksort(right)...)
	}

	return arr
}

func SelectPivot(arr []int) int {
	pivotIdx := 0
	if len(arr) < 3 {return pivotIdx}

	if arr[0] > arr[1] {
		if arr[0] < arr[2] {
			pivotIdx = 0
		} else if arr[1] > arr[2] {
			pivotIdx = 1
		} else {
			pivotIdx = 2
		}
	} else {
		if arr[0] > arr[2] {
			pivotIdx = 0
		} else if arr[1] < arr[2] {
			pivotIdx = 1
		} else {
			pivotIdx = 2
		}
	}
	
	return pivotIdx
}

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		cmpIdx := i-1
		insertVal := arr[i]
		for ; cmpIdx >= 0 && arr[cmpIdx] > insertVal; cmpIdx-- {
			arr[cmpIdx+1] = arr[cmpIdx]
		}
		arr[cmpIdx+1] = insertVal
	}
	return arr
}