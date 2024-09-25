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
		switch {
		case array[mid] == target:
			return mid
		case array[mid] > target:
			high = mid - 1
		case array[mid] < target:
			low = mid + 1
		}
	}
	return -1
}
