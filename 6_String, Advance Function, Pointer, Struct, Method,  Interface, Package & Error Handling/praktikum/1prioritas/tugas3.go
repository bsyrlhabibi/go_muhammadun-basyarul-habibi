package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	// your code here
	var result string

	for i := 0; i < len(a); i++ {
		for j := len(a); j > i; j-- {
			substring := a[i:j]
			if len(substring) > len(result) && strings.Index(b, substring) >= 0 {
				result = substring
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     //AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  //KANG
	fmt.Println(Compare("KI", "KIJANG"))      //KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) //KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    //ILA
}
