package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		if sum > 50 {
			break
		}
		sum += i
	}

	fmt.Println(sum) // 45
}
