package main

import (
	"bufio"
	"fmt"
	"os"
)

func ubahHuruf(sentence string) string {
	var pengganti string

	for i := 0; i < len(sentence); i++ {

		x := sentence[i]
		if x == 32 {
			pengganti = pengganti + string(x)
			// fmt.Println(string(pengganti))
			continue
		}
		if x > 90 {
			x = x - 32
		}
		x = x + 10
		if x > 90 {
			x = x - 26
		}
		pengganti = pengganti + string(x)
		// fmt.Println(string(pengganti))
	}
	return pengganti
}

func main() {
	fmt.Println("Ubah Kata")
	fmt.Printf("\nMasukkan Kata : ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading

	kata := scanner.Text()

	kata = ubahHuruf(kata)

	fmt.Println(string(kata))

}
