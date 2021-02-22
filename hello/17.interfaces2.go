package main

import "fmt"

func main() {
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 12, 44}
	fmt.Println(m) // map[age:18 name:Tom scores:[98 12 44]]
}
