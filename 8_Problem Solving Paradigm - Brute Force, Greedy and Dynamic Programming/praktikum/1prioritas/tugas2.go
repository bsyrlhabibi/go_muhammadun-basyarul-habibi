package main

import "fmt"

func pascal(numRows int) [][]int {
	segitigaPascal := make([][]int, numRows)
	for i := range segitigaPascal {
		segitigaPascal[i] = make([]int, i+1)
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				segitigaPascal[i][j] = 1
			} else {
				segitigaPascal[i][j] = segitigaPascal[i-1][j-1] + segitigaPascal[i-1][j]
			}
		}
	}

	return segitigaPascal
}

func main() {
	segitigaPascal := pascal(5)
	fmt.Println(segitigaPascal) // [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
