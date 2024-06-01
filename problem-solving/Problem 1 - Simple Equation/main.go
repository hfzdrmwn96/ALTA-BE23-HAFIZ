package main

import "fmt"

func SimpleEquations(a, b, c int) {
	found := false
	var x, y, z int
	for x = 1; x <= a/2; x++ {
		y = (a - x) / 2
		z = a - x - y

		if x*y*z == b && x*x+y*y+z*z == c {
			found = true
			break
		}

	}

	fmt.Printf("Input: %d %d %d\n", a, b, c)

	if found {
		fmt.Printf("Output: %d %d %d\n\n", x, y, z)
	} else {
		fmt.Println("Output: No solution\n")
	}

}

func main() {
	SimpleEquations(1, 2, 3)  // no solution
	SimpleEquations(6, 6, 14) // 1 2 3
	SimpleEquations(6, 8, 12) // 2 2 2
	SimpleEquations(4, 2, 13) // 2 2 2
}
