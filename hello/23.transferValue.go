// 传值和传指针
// 当我们传一个参数值到被调用函数里面时，实际上是传了这个值的一个copy。

package main

import "fmt"

// 简单的实现参数加1的操作
func add1(a int) int {
	a = a + 1
	return a
}

func main() {
	x := 3

	fmt.Println("x = ", x) // x =  3

	x1 := add1(x)

	fmt.Println("x + 1 = ", x1) // x + 1 =  4
	fmt.Println("x = ", x)      // x =  3
}
