package main

import "fmt"

type testInt func(int) bool // 声明一个函数类型

// 奇数
func isOdd(interger int) bool {
	if interger%2 == 0 {
		return false
	}
	return true
}

// 偶数
func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

// 声明的函数类型在这个地方当做一个参数
func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	fmt.Println("Odd elements of slice are:", odd) // Odd elements of slice are: [1 3 5 7]
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are:", even) // Even elements of slice are: [2 4]
}
