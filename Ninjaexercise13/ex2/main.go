package main

import (
	"bharathi/Ninjaexercise13/ex2/quote"
	"fmt"

	"bharathi/Ninjaexercise13/ex2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
