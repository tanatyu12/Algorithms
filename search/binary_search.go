package main

import (
	"fmt"
	"./utils"
	"time"
)

// best: O(1)
// average: O(log n)
// worst: O(log n)
func main() {
	n, target, err := utils.GetCommandArg()
	if(err != nil) {
		fmt.Println(err)
		return
	}
	// arr is sorted collection (binary search need sorted)
	arr := utils.MakeCollection(n)

	start := time.Now()
	
	// binary search
	low := 0
	high := n
	for low <= high {
		mid := (low + high) / 2
		if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		} else {
			fmt.Println("found")
			end := time.Now()
			fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
			return
		}
	}
	fmt.Println("not found")

	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}