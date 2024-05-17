package main

import "fmt"

func main() {
	var hd int
	var ha int
	var d int
	var b int

	fmt.Print("Masukkan Harga Barang : ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan diskon (%) : ")
	fmt.Scanln(&d)

	hd = (d * b) / 100
	ha = b - hd

	fmt.Print("Harga Akhir = Rp", ha)
}
