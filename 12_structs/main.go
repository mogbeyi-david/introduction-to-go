package main

import "fmt"
import "strconv"

type Person struct {
	firstName string
	lastName  string
}

func (person Person) greet() string {
	return "My name is " + person.firstName + strconv.Itoa(12)
}

func main() {
	var firstPerson = Person{"David", "Owumi"}
	fmt.Println(firstPerson.greet())
}
