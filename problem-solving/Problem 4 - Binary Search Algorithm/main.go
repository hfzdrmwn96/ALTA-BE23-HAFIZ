package main

import "fmt"

func BinarySearch(array []int, x int) {
	//code
	var result int = -1
	var beg, end int = 0, len(array) - 1
	for beg <= end {

		mid := (beg + end) / 2
		if array[mid] == x {
			result = mid
			break
		} else if array[mid] > x {
			end = mid - 1
		} else if array[mid] < x {
			beg = mid + 1
		}
	}

	fmt.Println(result)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  //2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 //3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  //6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) //-1
}
