package main

import "fmt"

func main() {
	fmt.Println("faktor_bilangan_2")
	var nilai int
	fmt.Printf("Masukkan angka : ")
	fmt.Scanln(&nilai)

	for i := nilai; i >= 1; i-- {
		if nilai%i == 0 {
			fmt.Println(i)
		}
	}
}
