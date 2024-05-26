package main

import (
	"fmt"
)

func ArrayUnique(arrayA, arrayB []int) []int {
	var hasil []int

	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			if arrayA[i] == arrayB[j] {
				break
			} else if j == len(arrayB)-1 {
				hasil = append(hasil, arrayA[i])
			}
		}
	}

	return hasil
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 18}))
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59}))
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))

}
