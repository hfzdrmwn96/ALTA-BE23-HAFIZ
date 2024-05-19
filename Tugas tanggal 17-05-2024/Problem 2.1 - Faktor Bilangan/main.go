package main

import "fmt"

func main() {

	fmt.Println("faktor_bilangan")
	var nilai int
	fmt.Printf("Masukkan angka : ")
	fmt.Scanln(&nilai)

	for i := 1; i <= nilai; i++ {
		if nilai%i == 0 {
			fmt.Println(i)
		}
	}
}
