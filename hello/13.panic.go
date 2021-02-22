package main

import "fmt"

func get(index int) int {
	arr := [3]int{2, 4, 3}
	return arr[index]
}

func main() {
	fmt.Println(get(5))
	fmt.Println("finished")
}

// panic: runtime error: index out of range [5] with length 3

// goroutine 1 [running]:
