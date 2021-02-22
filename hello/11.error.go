package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err) // open filename.txt: no such file or directory
	}
}
