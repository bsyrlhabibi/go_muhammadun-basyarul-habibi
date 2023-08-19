package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func palindrome(value string) bool {
	for i := 0; i < len(value)/2; i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Apakah Palindrome?")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan kata: ")
	scanner.Scan()
	input := scanner.Text()
	word := strings.TrimSpace(input)

	fmt.Printf("Captured: %s\n", word)

	result := palindrome(word)

	if result {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Bukan Palindrome")
	}
}
