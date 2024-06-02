package main

import "fmt"

func SimpleEquations(a, b, c int) {
	found := false
	var x, y, z int
	for x = 1; x <= a/3; x++ {
		for y = a - x - 1; y >= a/3; y-- {
			z = a - x - y
			if x*y*z == b && x*x+y*y+z*z == c {
				found = true
				break
			}

		}
		if found {
			break
		}

	}

	fmt.Printf("Input: %d %d %d\n", a, b, c)

	if found {
		fmt.Printf("Output: %d %d %d\n\n", x, y, z)
	} else {
		fmt.Printf("Output: No solution\n\n")
	}

}

func main() {
	SimpleEquations(1, 2, 3)         // no solution
	SimpleEquations(6, 6, 14)        // 1 2 3
	SimpleEquations(66, 10648, 1452) // 22 22 22
	SimpleEquations(4, 2, 13)        // no solution
	SimpleEquations(21, 231, 179)    // 3 11 7
	SimpleEquations(42, 1122, 782)   // 3 11 7
}
