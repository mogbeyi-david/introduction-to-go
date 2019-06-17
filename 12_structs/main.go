package main

import "fmt"
import "strconv"

type Person struct {
	firstName string
	lastName  string
	age int
	gender string
}

func (person Person) greet() string {
	return "My name is " + person.firstName + strconv.Itoa(12)
}

func (person *Person) hasBirthday()  {
	person.age++
}

func (person *Person) getsMarried(spouseLastname string)  {
	if person.gender == "m" {
		return
	} else{
		person.lastName = spouseLastname
	}
}

func main() {
	var firstPerson = Person{"David", "Owumi", 32, "m"}
	firstPerson.hasBirthday()
	firstPerson.hasBirthday()
	firstPerson.hasBirthday()
	firstPerson.hasBirthday()
	firstPerson.getsMarried("Williams")
	fmt.Println(firstPerson)
}
