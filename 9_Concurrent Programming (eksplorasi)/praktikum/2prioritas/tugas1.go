package main

import (
	"fmt"
	"sync"
)

func frekuensi(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	return m
}

func frekuensiParalel(teks []string) map[rune]int {
	var wg sync.WaitGroup
	hasil := make(map[rune]int)
	kunciMutex := sync.Mutex{}

	for _, teksSatu := range teks {
		wg.Add(1)
		go func(teksSatu string) {
			defer wg.Done()
			hasilTeks := frekuensi(teksSatu)
			kunciMutex.Lock()
			defer kunciMutex.Unlock()
			for k, v := range hasilTeks {
				hasil[k] += v
			}
		}(teksSatu)
	}

	wg.Wait()
	return hasil
}

func main() {
	teks := []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
	}

	hasil := frekuensiParalel(teks)

	for k, v := range hasil {
		fmt.Printf("%c: %d\n", k, v)
	}
}
