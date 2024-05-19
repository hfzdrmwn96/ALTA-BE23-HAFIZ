package main

import (
	"fmt"
)

func bilanganprima(bilangan int) bool {
	var adalahprima bool
	if bilangan < 2 {
		fmt.Println(bilangan, "adalah bukan bilangan prima")
	} else {
		for i := 2; i <= bilangan; i++ {
			if i == bilangan && i > 1 {
				fmt.Println(bilangan, "adalah bilangan prima")
				adalahprima = true
				break
			}
			if bilangan%i == 0 {
				fmt.Println(bilangan, "adalah bukan bilangan prima")
				adalahprima = false
				break
			}
		}
	}
	return adalahprima
}

func main() {
	fmt.Println("Full Prima")
	var nilai int
	var fullprima bool
	fmt.Printf("Masukkan nilai = ")
	fmt.Scanln(&nilai)
	fmt.Println("nilai =", nilai)
	fullprima = bilanganprima(nilai)

	if !fullprima {
		fmt.Println(nilai, "adalah bukan bilangan full prima")
	} else {
		var nilai_cek int
		var nilai_sisa = nilai

		for nilai_sisa > 0 {
			nilai_cek = nilai_sisa % 10
			fullprima = bilanganprima(nilai_cek)
			if !fullprima {
				fmt.Println(nilai, "adalah bukan bilangan full prima")
				break
			}
			nilai_sisa = (nilai_sisa - nilai_cek) / 10
		}
		if fullprima {
			fmt.Println(nilai, "adalah bilangan full prima")
		}
	}

}
