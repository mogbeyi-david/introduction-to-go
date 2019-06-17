package main

import "fmt"

func main() {
	var name = "David"
	var age int32 = 37
	const isCool = true
	var size float32 = 1.3;
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", size)
}
