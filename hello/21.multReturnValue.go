package main

import "fmt"

func SumAndProduct(A, B int) (add int, Meultiplied int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4
	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}

// 3 + 4 = 7
// 3 * 4 = 12
