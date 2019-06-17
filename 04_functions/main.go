package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func main() {
	fmt.Println(greeting("David"))
	fmt.Println(getSum(1, 1))
}
