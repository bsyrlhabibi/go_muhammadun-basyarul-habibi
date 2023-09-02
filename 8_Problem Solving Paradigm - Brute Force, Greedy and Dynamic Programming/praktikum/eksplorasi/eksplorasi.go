package main

import "fmt"

func angkaRomawi(n int) string {
	romawi := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	angka := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	result := ""

	for i := 0; i < len(romawi); i++ {
		for n >= angka[i] {
			result += romawi[i]
			n -= angka[i]
		}
	}

	return result
}

func main() {
	fmt.Println(angkaRomawi(4))    // IV
	fmt.Println(angkaRomawi(9))    // IX
	fmt.Println(angkaRomawi(23))   // XXIII
	fmt.Println(angkaRomawi(2021)) // MMXXI
	fmt.Println(angkaRomawi(1646)) // MDCXLVI
}
