package main

import "fmt"

func getMinMax(numbers ...*int) (min, max int) {
	// your code here
	min = *numbers[0]
	max = *numbers[0]

	for i := 0; i < len(numbers); i++ {
		number := *numbers[i]
		if number > max {
			max = number
		} else if number < min {
			min = number
		}
	}

	return min, max

}

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Println("Input:")
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	fmt.Println("Output:")

	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai Min : ", min)
	fmt.Println("Nilai Max : ", max)
}
