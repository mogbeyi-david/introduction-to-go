package main

import "fmt"

func main() {
	//ids := []int{1, 2, 3, 4, 5, 6, 7}
	//var sum = 0
	//for _, id := range ids {
	//	sum += id
	//}
	emails := map[string]string{"Bob": "bob@gmail.com", "Javascript": "Javascript@gmail.com"}
	for key, value:= range emails{
		fmt.Println(key, value)
	}
}
