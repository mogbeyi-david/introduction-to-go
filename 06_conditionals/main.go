package main

import "fmt"

func main() {
	const x = 15
	const y = 10

	if x > y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than or equal to y")
	}

	const color = "blue"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

}
