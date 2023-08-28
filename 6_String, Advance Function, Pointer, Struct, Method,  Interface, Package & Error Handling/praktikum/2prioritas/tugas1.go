package main

import "fmt"

func caesar(offset int, input string) string {
	// your code here
	result := ""

	for _, char := range input {

		if char >= 'a' {
			shifted := (int(char) - 'a' + offset) % 26
			result += string(rune(shifted) + 'a')
		} else if char == ' ' {
			result += " "
		}
	}

	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
