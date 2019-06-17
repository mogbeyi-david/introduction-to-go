package main

import "fmt"

func main() {
	// Arrays...
	//var fruitArray [2]string
	//fruitArray[0] = "Apple"
	//fruitArray[1] = "orange"
	//fmt.Println(fruitArray)
	//fruitArray := [2]string{"Apple", "Orange"}
	//fmt.Println(fruitArray)

	fruitSlice := []string{"apple", "mango", "orange", "grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
