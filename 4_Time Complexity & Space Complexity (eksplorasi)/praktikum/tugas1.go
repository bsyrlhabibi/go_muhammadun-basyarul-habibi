package main

import (
	"fmt"
	"math"
)

func primeNumber(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	inputs := []int{1000000007, 1500450271, 35}

	for _, number := range inputs {
		if primeNumber(number) {
			fmt.Printf("Input: %d\nOutput: Bilangan Prima\n\n", number)
		} else {
			fmt.Printf("Input: %d\nOutput: Bukan Bilangan Prima\n\n", number)
		}
	}

	fmt.Println(primeNumber(1000000007)) // true
	fmt.Println(primeNumber(13))         // true
	fmt.Println(primeNumber(17))         // true
	fmt.Println(primeNumber(20))         // false
	fmt.Println(primeNumber(35))         // false
}
