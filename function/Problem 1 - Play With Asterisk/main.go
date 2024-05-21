package main

import "fmt"

func playWithAsterisk(n int) string {
	var pwa string
	for i := 0; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			// fmt.Printf(" ")
			pwa = pwa + " "
		}
		for k := 0; k < i; k++ {
			// fmt.Printf("* ")
			pwa = pwa + "* "
		}
		// fmt.Printf("\n")
		pwa = pwa + "\n"
	}
	return pwa
}

func main() {
	fmt.Println("Problem 1 - Playing With Asterisk")
	var input int
	var output string
	fmt.Printf("Masukkan input angka : ")
	fmt.Scanln(&input)
	output = playWithAsterisk(input)
	fmt.Printf("%s", output)
}
