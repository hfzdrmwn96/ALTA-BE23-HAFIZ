package main

import (
	"fmt"
)

func cetakTabelPerkalian(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d", j*i)
			if j*i < 10 {
				fmt.Printf("  ")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")

	}
}

func main() {
	fmt.Println("Problem 3 - Cetak Tabel Perkalian")
	var input int
	//var output string
	fmt.Printf("Masukkan input angka : ")
	fmt.Scanln(&input)
	cetakTabelPerkalian(input)

}
