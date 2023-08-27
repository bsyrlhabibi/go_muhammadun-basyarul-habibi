package main

import (
	"fmt"
	"math"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}

	leftToRight := 0
	rightToLeft := 0

	n := len(matrix)

	for i := 0; i < n; i++ {
		leftToRight += matrix[i][i]
		rightToLeft += matrix[i][n-i-1]
	}

	selisihAbsolut := int(math.Abs(float64(leftToRight - rightToLeft)))

	fmt.Println("Diagonal kiri ke kanan =", leftToRight)
	fmt.Println("Diagonal kanan ke kiri =", rightToLeft)
	fmt.Println("Perbedaan mutlak =", selisihAbsolut)
}
