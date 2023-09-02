package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	// your code here
	for _, card := range cards {
		for _, number := range deck {
			if card[0] == number || card[1] == number {
				return card
			}
		}
	}

	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6, 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))              // "Tutup kartu"
}
