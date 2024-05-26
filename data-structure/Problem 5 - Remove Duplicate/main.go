package main

import (
	"fmt"
)

func RemoveDuplicates(array []int) int {
	var hasil []int

	for i := 0; i < len(array); i++ {
		if i == len(array)-1 || array[i] != array[i+1] {
			hasil = append(hasil, array[i])
		}

	}

	return len(hasil)
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 6, 9, 9}))
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9}))
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))
}
