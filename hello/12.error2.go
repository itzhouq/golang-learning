package main

import (
	"errors"
	"fmt"
)

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello, ", name)
	return nil
}

func main() {
	if err := hello(""); err != nil {
		fmt.Println(err) // error: name is null
	}
}
