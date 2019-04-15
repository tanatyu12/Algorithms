package main

import "fmt"

// best: O(n)
// average: O(n^2)
// worst: O(n^2)
func main() {
	arr := []int{4, 6, 1, 12, 9, 10, 3, 2, 7, 13, 15, 8, 5, 11, 14}

	for i := 1; i < len(arr); i++ {
		cmpIdx := i-1
		insertVal := arr[i]
		for ; cmpIdx >= 0 && arr[cmpIdx] > insertVal; cmpIdx-- {
			arr[cmpIdx+1] = arr[cmpIdx]
		}
		arr[cmpIdx+1] = insertVal
	}
	fmt.Println(arr)
}

