package main

import "fmt"

func main() {
	var T float64
	var r float64
	var Lp float64

	fmt.Print("Masukkan nilai T : ")
	fmt.Scanln(&T)
	fmt.Print("Masukkan nilai r : ")
	fmt.Scanln(&r)

	Lp = 2 * 3.14 * r * (r + T)

	fmt.Printf("Luas Permukaan Tabung =  %.2f", Lp)
}
