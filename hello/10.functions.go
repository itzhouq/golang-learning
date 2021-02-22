package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}

func main() {
	quo, rem := div(100, 17)
	fmt.Println(quo, rem)      // 5, 15
	fmt.Println(add(100, 17))  // 117
	fmt.Println(add2(200, 12)) // 212
}
