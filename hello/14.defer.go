package main

import "fmt"

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happend!    ", r)
			ret = -1
		}
	}()

	arr := [3]int{2, 4, 3}
	return arr[index]
}

func main() {
	fmt.Println(get(5))
	fmt.Println("finished")
}

// Some error happend!     runtime error: index out of range [5] with length 3
// -1
// finished
