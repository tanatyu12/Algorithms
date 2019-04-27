package main

import "fmt"


// best: O(n)
// average: O(n)
// worst: O(n)

type Entry struct {
	value int
	next *Entry
}

type Bucket struct {
	size int
	head *Entry
}

/**
  hash(arr[i]) and bucketNum are decided by character of elements.
  if multiple elements exist in bucket, extract sorted elements. (use insertion sort)
**/
func main() {
	arr := []int{4, 6, 1, 12, 9, 10, 3, 2, 7, 13, 15, 8, 5, 11, 14}

	bucketNum := 8
	buckets := make([]Bucket, bucketNum)
	for i := 0; i < len(arr); i++ {
		k := hash(arr[i])
		
		e := Entry{arr[i], nil}
		if buckets[k].head == nil {
			buckets[k].head = &e
		} else {
			e.next = buckets[k].head
			buckets[k].head = &e
		}
		buckets[k].size++
	}

	sortedArr := Extract(&buckets)
	fmt.Println(sortedArr)
}

func hash(x int) int {
	return x / 2
}

func InsertionSort(bucket Bucket, arrPtr *[]int) {
	arr := make([]int, bucket.size)
	e := bucket.head
	for i := 0; i < bucket.size; i++ {
		arr[i] = e.value
		e = e.next
	}
	for i := 1; i < len(arr); i++ {
		cmpIdx := i-1
		insertVal := arr[i]
		for ; cmpIdx >= 0 && arr[cmpIdx] > insertVal; cmpIdx-- {
			arr[cmpIdx+1] = arr[cmpIdx]
		}
		arr[cmpIdx+1] = insertVal
	}
	*arrPtr = append(*arrPtr, arr...)
}

func Extract(bucketsPtr *[]Bucket) []int {
	arr := make([]int, 0)
	idx := 0
	for i := 0; i < len(*bucketsPtr); i++ {
		if (*bucketsPtr)[i].size == 0 {
			continue
		} else if (*bucketsPtr)[i].size == 1 {
			arr = append(arr, (*bucketsPtr)[i].head.value)
			idx++
		} else {
			InsertionSort((*bucketsPtr)[i], &arr)
			idx += (*bucketsPtr)[i].size
		}
	}
	return arr
}

