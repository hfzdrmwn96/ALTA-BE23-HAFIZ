package main

import (
	"fmt"
)

func caesar(offset int, input string) string {
	var hasil string

	offset = offset % 26
	for i := 0; i < len(input); i++ {
		if input[i]+byte(offset) > 122 {
			hasil = hasil + string(input[i]+byte(offset)-byte(26))
		} else {
			hasil = hasil + string(input[i]+byte(offset))
		}
	}

	return hasil
}

func main() {
	fmt.Println(caesar(3, "abc"))
	fmt.Println(caesar(2, "alta"))
	fmt.Println(caesar(3, "alterraacademy"))
	fmt.Println(caesar(10, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

}
