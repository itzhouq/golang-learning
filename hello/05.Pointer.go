package main

import "fmt"

func main() {
	/* 类型定义时使用符号*， 对于一个已经存在的变量，使用&获取变量的地址 */
	str := "Golang"
	var p *string = &str // p是指向str的指针
	fmt.Println(str, p)  // Golang 0xc000108040

	*p = "Hello"
	fmt.Println(str, p) // Hello 0xc00008e1e0
	fmt.Println("========")
	num := 100
	add(num)
	fmt.Println(num) // 100
	realAdd(&num)
	fmt.Println(num) // 101, 指针传递，num被修改

}

func add(num int) {
	num++
}

func realAdd(num *int) {
	*num++
}
