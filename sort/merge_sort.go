package main

import "fmt"

// best: O(n log n)
// average: O(n log n)
// worst: O(n log n)
func main() {
	arr := []int{4, 6, 1, 12, 9, 10, 3, 2, 7, 13, 15, 8, 5, 11, 14}
	copy := append([]int{}, arr...)
	MergeSort(&arr, &copy, 0, len(arr))
	fmt.Println(copy)
}

func MergeSort(arrp *[]int, result *[]int, start int, end int) {
	if end - start < 2 {
		return
	}
	if end - start == 2 {
		if (*result)[start] > (*result)[start+1] {
			tmp := (*result)[start]
			(*result)[start] = (*result)[start+1]
			(*result)[start+1] = tmp
		}
		return
	}

	mid := (start + end) / 2
	MergeSort(result, arrp, start, mid)
	MergeSort(result, arrp, mid, end)

	i := start
	j := mid
	for idx := start; idx < end; idx++ {
		if j >= end || (i < mid && (*arrp)[i] < (*arrp)[j]) {
			(*result)[idx] = (*arrp)[i]
			i++
		} else {
			(*result)[idx] = (*arrp)[j]
			j++
		}
	}
}