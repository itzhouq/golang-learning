// 变参
package main

import "fmt"

func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	myfunc(2, 3, 5, 7)
}

// And the number is: 2
// And the number is: 3
// And the number is: 5
// And the number is: 7
