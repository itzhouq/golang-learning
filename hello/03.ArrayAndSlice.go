package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	fmt.Println(arr)
	fmt.Println("=====")

	slice1 := make([]float32, 0)          // 长度为0的切片
	slice2 := make([]float32, 3, 5)       // 长度为3容量为5的切片
	fmt.Println(len(slice1), cap(slice1)) // 0 0
	fmt.Println(len(slice2), cap(slice2)) // 3 5

	fmt.Println("======")

	slice2 = append(slice2, 1, 2, 3, 4)
	fmt.Println(slice2)                   // [0 0 0 1 2 3 4]
	fmt.Println(len(slice2), cap(slice2)) // 7 12

	// 子切片 [start, end)
	sub1 := slice2[3:]
	sub2 := slice2[:3]
	sub3 := slice2[1:4]
	fmt.Println(sub1) // [1 2 3 4]
	fmt.Println(sub2) // [0 0 0]
	fmt.Println(sub3) // [0 0 1]

	// 合并切片
	combined := append(sub1, sub2...) // [1 2 3 4 0 0 0]
	fmt.Println(combined)

}
