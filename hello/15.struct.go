package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (stu *Student) hello(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func main() {
	stu := &Student{
		name: "Tom",
	}

	msg := stu.hello("Jack")
	fmt.Println(msg) // hello Jack, I am Tom
}
