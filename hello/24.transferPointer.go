// 传递指针
// &x 取出x的指针
// *int 指针类型
// 只有add1函数知道x变量所在的地址，才能修改x的值。所以需要将x所在的地址&x传入函数
// 此时参数仍然是按copy传递的，只是copy的值一个指针
package main

import "fmt"

// 简单的实现参数+1的操作
func add1(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	x := 3

	fmt.Println("x = ", x) // x =  3

	x1 := add1(&x)

	fmt.Println("x + 1 = ", x1) // x + 1 =  4
	fmt.Println("x = ", x)      // x =  4
}
