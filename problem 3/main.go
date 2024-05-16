package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Masukkan nama : ")
	fmt.Scanln(&nama)
	fmt.Print("Hello ", nama, "! Saya Golang bahasa yang sangat menyenangkan")
}
