package main

import "fmt"

func getMinMax(numbers ...*int) (min int, max int) {

	for i := 0; i < len(numbers); i++ {
		if i == 0 {
			min = *numbers[i]
			max = *numbers[i]
		}
		if min > *numbers[i] {
			min = *numbers[i]
		}
		if max < *numbers[i] {
			max = *numbers[i]
		}

	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Printf("masukkan angka ke - 1 : ")
	fmt.Scan(&a1)
	fmt.Printf("masukkan angka ke - 2 : ")
	fmt.Scan(&a2)
	fmt.Printf("masukkan angka ke - 3 : ")
	fmt.Scan(&a3)
	fmt.Printf("masukkan angka ke - 4 : ")
	fmt.Scan(&a4)
	fmt.Printf("masukkan angka ke - 5 : ")
	fmt.Scan(&a5)
	fmt.Printf("masukkan angka ke - 6 : ")
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min ", min)
	fmt.Println("Niali max ", max)
}
