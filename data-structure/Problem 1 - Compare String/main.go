package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var hasil string
	if len(a) < len(b) {
		temp := a
		a = b
		b = temp
	}
	if strings.Contains(a, b) {
		hasil = b
	} else {
		hasil = "False"
	}

	return hasil
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
	//fmt.Println(Compare("ILALANG", "LALAN"))  // LANG
}
