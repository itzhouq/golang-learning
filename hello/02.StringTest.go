package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "Golang"
	str2 := "Go语言"
	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // uint8
	fmt.Println(str1[2], string(str1[2]))       // 108 l
	fmt.Printf("%d %c\n", str2[2], str2[2])     // 232 è
	fmt.Println("len(str2): ", len(str2))       // len(str2):  8
	fmt.Println("=========")

	runeArr := []rune(str2)
	fmt.Println(reflect.TypeOf(runeArr[2]).Kind()) // int32
	fmt.Println(runeArr[2], string(runeArr[2]))    // 35821 语
	fmt.Println("len(runeArr): ", len(runeArr))    // len(runeArr):  4
}
