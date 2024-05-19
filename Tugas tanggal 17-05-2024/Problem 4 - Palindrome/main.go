package main

import (
	"bufio"
	"fmt"
	"os"
	//"reflect"
)

func main() {
	fmt.Println("palindrome")
	fmt.Printf("\nMasukkan Kata : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading

	kata := scanner.Text()
	j := len(kata) - 1

	//fmt.Println("Jumlah huruf : ", j+1)
	// var reflectValue = reflect.ValueOf(j)

	// fmt.Println("tipe  variabel :", reflectValue.Type())

	//fmt.Println("Kata yang diinput : ", kata)

	for i := 0; i <= len(kata)-1; i++ {
		// fmt.Println(i, " ", string(kata[i]), " ", string(kata[j]))
		if j == 0 && kata[i] == kata[j] {
			fmt.Println("true")
			break
		} else if kata[i] != kata[j] {
			fmt.Println("false")
			break
		}

		j--

	}

}
