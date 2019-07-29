package main

import (
	"fmt"
	"./utils"
	"time"
)

// best: O(1)
// average: O(n)
// worst: O(n)
func main() {
	n, target, err := utils.GetCommandArg()
	if(err != nil) {
		fmt.Println(err)
		return
	}
	// arr is sorted collection
	arr := utils.MakeCollection(n)

	start := time.Now()

	// sequential search
	for _, val := range arr {
		if val == target {
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