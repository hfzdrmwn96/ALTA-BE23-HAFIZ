package main

import "fmt"

func main() {
	fmt.Println("bilangan_prima")
	var nilai int
	fmt.Printf("Masukkan angka : ")
	fmt.Scanln(&nilai)
	if nilai < 2 {
		fmt.Println("false")
	} else {
		for i := 2; i <= nilai; i++ {
			if i == nilai && i > 1 {
				fmt.Println("true")
				break
			}
			if nilai%i == 0 {
				fmt.Println("false")
				break
			}
		}
	}
}
