package main

import (
	"fmt"
)

func ArrayUnique(k int, arr []int) int {
	var hasil int = 0

	for i := 0; i < len(arr)-k; i++ {
		pembanding := 0
		for j := i; j < i+k; j++ {
			pembanding = pembanding + arr[j]
		}
		if hasil < pembanding {
			hasil = pembanding
		}
	}

	return hasil
}

func main() {
	fmt.Println(ArrayUnique(3, []int{2, 1, 5, 1, 3, 2}))
	fmt.Println(ArrayUnique(2, []int{2, 3, 4, 1, 5}))
	fmt.Println(ArrayUnique(2, []int{2, 1, 4, 1, 1}))
	fmt.Println(ArrayUnique(3, []int{2, 1, 4, 1, 1}))
	fmt.Println(ArrayUnique(4, []int{2, 1, 4, 1, 1}))

}
