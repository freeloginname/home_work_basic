package binarySearch

import (
	"fmt"
)

func BinarySearch(array []int, target int) int {

	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		fmt.Printf("low = %d, high = %d, mid = %d\n", low, high, mid)
		if array[mid] == target {
			return mid
		} else if array[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5

	index := BinarySearch(items, target)
	fmt.Printf("Found %d at index %d\n", target, index)
}
