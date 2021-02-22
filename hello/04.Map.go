package main

import "fmt"

func main() {
	// 仅声明
	m1 := make(map[string]int)

	// 声明时赋值
	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	// 修改/赋值
	m1["Tom"] = 18
	m2["Sam"] = "Male1111"

	fmt.Println(m1) // map[Tom:18]
	fmt.Println(m2) // map[Alice:Female Sam:Male1111]
}
