package main

import "fmt"

func main() {

	fmt.Println("konversi_nilai")
	var nilai int
	fmt.Printf("Masukkan nilai siswa : ")
	fmt.Scanln(&nilai)
	if nilai >= 0 && nilai <= 100 {
		if nilai >= 80 {
			fmt.Println("Nilai huruf : A")
		} else if nilai >= 65 {
			fmt.Println("Nilai huruf : B+")
		} else if nilai >= 50 {
			fmt.Println("Nilai huruf : B")
		} else if nilai >= 35 {
			fmt.Println("Nilai huruf : C")
		} else if nilai >= 0 {
			fmt.Println("Nilai huruf : D")
		}
	} else {
		fmt.Println("Pastikan memasukkan angka antara 0 - 100")
	}
}
