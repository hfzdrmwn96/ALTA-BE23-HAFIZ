package main

import "fmt"

func main() {
	var alas float64
	var tinggi float64
	var luas float64

	fmt.Print("Masukkan nilai alas : ")
	fmt.Scanln(&alas)
	fmt.Print("Masukkan nilai tinggi : ")
	fmt.Scanln(&tinggi)

	luas = alas * tinggi / 2

	fmt.Printf("Luas Segitiga = %.2f", luas)
}
