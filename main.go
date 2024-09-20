package main

import (
	"fmt"
)

const (
	A int = 1
	B int = 2
	c int = 3
)

func main() {
	var num1 int32 = 120
	var num2 int16 = 125
	fmt.Println(num1, num2)
	fmt.Println(A, B, c)
	fmt.Println("hello world")
	var a, b = bio(20, "joy")
	fmt.Println(b)
	fmt.Printf("you are %v years old\n", a)

	var _, result = bio(10, "mahmud")
	fmt.Println((result))

}
func bio(age int, name string) (myage int, result string) {
	result = "hello " + name
	myage = age + 10
	return
}
