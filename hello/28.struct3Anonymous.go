package main

import "fmt"

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      // 匿名字段，struct
	Skills     // 匿名字段，自定义的类型string slice
	int        // 内置类型作为匿名字段
	speciality string
}

func main() {
	// 初始化学生Jane
	jane := Student{Human: Human{"Jane", 35, 100}, speciality: "Biology"}
	// 现在我们来访问相应的字段
	fmt.Println("Her name is ", jane.name)             // Her name is  Jane
	fmt.Println("Her age is ", jane.age)               // Her age is  35
	fmt.Println("Her weight is ", jane.weight)         // Her weight is  100
	fmt.Println("Her speciality is ", jane.speciality) // Her speciality is  Biology
	// 我们来修改他的skill技能字段
	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills) /// Her skills are  [anatomy]
	fmt.Println("She acquired two new ones ")   // She acquired two new ones
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills) // Her skills now are  [anatomy physics golang]
	// 修改匿名内置类型字段
	jane.int = 3
	fmt.Println("Her preferred number is", jane.int) // Her preferred number is 3
}
