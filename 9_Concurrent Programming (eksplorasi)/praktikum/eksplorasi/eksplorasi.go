package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	fmt.Println("products data")
	fmt.Println("===")

	apiURL := "https://fakestoreapi.com/products"

	productsChannel := make(chan []Product)

	var wg sync.WaitGroup

	numberGoroutines := 5

	for i := 0; i < numberGoroutines; i++ {
		wg.Add(1)
		go fetchProducts(apiURL, productsChannel, &wg)
	}

	go func() {
		wg.Wait()
		close(productsChannel)
	}()

	for products := range productsChannel {
		for _, product := range products {
			fmt.Printf("title: %s\n", product.Title)
			fmt.Printf("price: %.2f\n", product.Price)
			fmt.Printf("category: %s\n", product.Category)
			fmt.Println("===")
			fmt.Println("===")
		}
	}
}

func fetchProducts(apiURL string, productsChannel chan []Product, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var products []struct {
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Category string  `json:"category"`
	}

	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var productData []Product
	for _, p := range products {
		productData = append(productData, Product{
			Title:    p.Title,
			Price:    p.Price,
			Category: p.Category,
		})
	}

	productsChannel <- productData
}
