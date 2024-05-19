package main

import "fmt"

func main() {
	fmt.Println("exponentiation")
	var x, n, hasil int

	fmt.Printf("Masukkan nilai x : ")
	fmt.Scanln(&x)
	fmt.Printf("Masukkan nilai n : ")
	fmt.Scanln(&n)

	hasil = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * x
	}

	fmt.Printf("Hasil eksponensiasi : %d", hasil)

}
