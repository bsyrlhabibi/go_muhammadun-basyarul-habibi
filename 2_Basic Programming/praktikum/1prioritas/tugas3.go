package main

import "fmt"

func main() {
	var value int
	fmt.Print("Input Value: ")
	fmt.Scanln(&value)

	if value < 0 || value > 100 {
		println("Nilai Invalid")
	} else if value >= 80 {
		println("Grade: A")
	} else if value >= 65 {
		println("Grade: B")
	} else if value >= 50 {
		println("Grade: C")
	} else if value >= 35 {
		println("Grade: D")
	} else {
		println("Grade: E")
	}
}
