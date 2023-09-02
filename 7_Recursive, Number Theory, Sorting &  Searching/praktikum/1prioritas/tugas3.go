package main

import "fmt"

func primeX(number int) int {
	// your code here
	if number <= 0 {
		return 0
	}

	count := 0
	current := 2

	for {
		isPrime := true
		for i := 2; i*i <= current; i++ {
			if current%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
		if count == number {
			return current
		}
		current++
	}
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
