package main

import "fmt"

func pow(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		temp := pow(x, n/2)
		return temp * temp
	} else {
		temp := pow(x, (n-1)/2)
		return x * temp * temp
	}
}

func main() {
	x1, n1 := 2, 3
	result1 := pow(x1, n1)
	fmt.Printf("Input: x = %d, n = %d\n", x1, n1)
	fmt.Printf("Output: %d\n\n", result1)

	x2, n2 := 7, 2
	result2 := pow(x2, n2)
	fmt.Printf("Input: x = %d, n = %d\n", x2, n2)
	fmt.Printf("Output: %d\n\n", result2)

	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(5, 3))  // 125
	fmt.Println(pow(10, 2)) // 100
	fmt.Println(pow(2, 5))  // 32
	fmt.Println(pow(7, 3))  // 343
}
