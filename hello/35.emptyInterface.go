package main

import "fmt"

// 定义a 为空接口
var a interface{}
var i int = 5

var s string = "Hello Word"

// a 可以存储任意类型的数值

func main() {

	fmt.Println("a = ", a) // a =  <nil>
	a = i
	fmt.Println("a = ", a) // a =  5
	a = s
	fmt.Println("a = ", a) // a =  Hello Word
}

// 一个函数把interface{}作为参数，
// 那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。
