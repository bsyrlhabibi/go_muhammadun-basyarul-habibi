package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	total := 0
	count := len(s.score)
	for _, i := range s.score {
		total += i
	}

	if count > 0 {
		return float64(total) / float64(count)
	}
	return 0
}

func (s Student) Min() (min int, name string) {
	if len(s.score) == 0 {
		return 0, ""
	}

	min = s.score[0]
	name = s.name[0]

	for i := 1; i < len(s.score); i++ {
		if s.score[i] < min {
			min = s.score[i]
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	if len(s.score) == 0 {
		return 0, ""
	}

	max = s.score[0]
	name = s.name[0]

	for i := 1; i < len(s.score); i++ {
		if s.score[i] > max {
			max = s.score[i]
			name = s.name[i]
		}
	}

	return max, name
}

func main() {
	var a = Student{}

	for i := 1; i <= 5; i++ {
		var name string
		fmt.Print("Input  " + string(i) + " Student’s Name :  ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input  " + name + " Score :  ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is ", a.Average())

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is :  "+nameMin+" (", scoreMin, ")")

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is :  "+nameMax+" (", scoreMax, ")")
}
