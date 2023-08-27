package main

import (
	"fmt"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	// your code here
	merge := make(map[string]bool)
	array := []string{}

	for _, element := range arrayA {
		if _, ada := merge[element]; !ada {
			merge[element] = true
			array = append(array, element)
		}
	}

	for _, element := range arrayB {
		if _, ada := merge[element]; !ada {
			merge[element] = true
			array = append(array, element)
		}
	}

	return array
}

func main() {
	// Test cases
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]
	fmt.Println(ArrayMerge([]string{}, []string{}))
	// []
}
