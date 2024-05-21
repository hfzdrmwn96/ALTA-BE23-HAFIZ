package main

import "fmt"

func DrawXYZ(n int) string {
	var pwa string
	for i := 1; i <= n*n; i++ {
		if i%3 == 0 {
			// fmt.Printf("X ")
			pwa = pwa + "X "
		} else if i%2 == 0 {
			// fmt.Printf("Z ")
			pwa = pwa + "Z "
		} else {
			// fmt.Printf("Y ")
			pwa = pwa + "Y "
		}

		if i%n == 0 {
			// fmt.Printf("\n")
			pwa = pwa + "\n"
		}
		// fmt.Printf(" ")
		// pwa = pwa + " "

	}
	return pwa
}

func main() {
	fmt.Println("Problem 2 - Draw XYZ")
	var input int
	//var output string
	fmt.Printf("Masukkan input angka : ")
	fmt.Scanln(&input)
	//output = playWithAsterisk(input)
	fmt.Println(DrawXYZ(input))
}
