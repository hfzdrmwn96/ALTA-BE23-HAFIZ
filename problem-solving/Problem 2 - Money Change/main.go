package main

import "fmt"

func moneyChange(money int) []int {
	var kembalian []int
	var nominal = [11]int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	for i := len(nominal) - 1; i >= 0; i-- {
		if money >= nominal[i] {
			money -= nominal[i]
			kembalian = append(kembalian, nominal[i])
			i++
		}
	}

	return kembalian
}

func main() {
	fmt.Println(moneyChange(123))   // [100 20 1 1 1]
	fmt.Println(moneyChange(432))   // [200 200 20 10 1 1]
	fmt.Println(moneyChange(543))   // [500 20 20 1 1 1]
	fmt.Println(moneyChange(7752))  // [5000 2000 500 200 50 1 1]
	fmt.Println(moneyChange(15321)) // [10000 5000 200 100 20 1]
}
