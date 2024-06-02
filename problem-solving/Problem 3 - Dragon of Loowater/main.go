package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	//code
	var result [2]int
	var isFall bool = false
	sort.Ints(knightHeight)
	sort.Ints(dragonHead)
	for _, dh := range dragonHead {
		result[0] += dh
		for i := 0; i < len(knightHeight); i++ {
			if dh <= knightHeight[i] {
				result[1] += knightHeight[i]
				knightHeight = append(knightHeight[:i], knightHeight[i+1:]...)
				// fmt.Println(knightHeight)
				break
			} else if i == len(knightHeight)-1 {
				isFall = true
			}
			if isFall {
				break
			}
		}
	}

	if !isFall && result[1] > result[0] {
		fmt.Println(result[1])
	} else {
		fmt.Println("knight fall")
	}
}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})                                        // 11
	DragonOfLoowater([]int{5, 10}, []int{5})                                             // knight fall
	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})                                     // knight fall
	DragonOfLoowater([]int{7, 2, 3, 4, 5, 6, 4}, []int{7, 5, 4, 2, 2, 1, 4, 3, 1, 2, 8}) // 33
}
