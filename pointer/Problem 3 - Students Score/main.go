package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var avg float64
	for i := 0; i < len(s.score); i++ {
		avg += float64(s.score[i])

	}
	avg /= float64(len(s.score))
	return avg
}

func (s Student) Min() (min int, name string) {

	for i := 0; i < len(s.score); i++ {
		if i == 0 || min > s.score[i] {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return
}

func (s Student) Max() (min int, name string) {
	for i := 0; i < len(s.score); i++ {
		if i == 0 || min < s.score[i] {
			min = s.score[i]
			name = s.name[i]
		}
	}
	return
}

func main() {
	var a Student

	for i := 1; i < 6; i++ {
		var name string
		fmt.Print("Input ", i, " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + "'s Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}
	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
}
