package main

import (
	"example/calc"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())   // Hello, world.
	fmt.Println(calc.Add(10, 3)) // 13
}
