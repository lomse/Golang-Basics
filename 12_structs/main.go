package main

import (
	"fmt"
	"strconv"
)

// Person is a cool struct
type Person struct {
	// firstName string
	// lastName  string
	// age       int
	// city      string
	// gender    string

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver or getter)
func (person Person) greet() string {
	return "hello my name is " + person.firstName + " and I'm " + strconv.Itoa(person.age) + " years old"
}

// updateCity is a pointer receiver or setter
func (person *Person) updateCity() {
	person.city = "Kumasi"
}

func main() {
	firstName := "lom"
	lastName := "Se"
	age := 25
	city := "Accra"
	gender := "m"

	person := Person{firstName, lastName, city, gender, age}

	fmt.Println(person)

	fmt.Println(person.firstName)

	fmt.Println(person.greet())

	person.updateCity()

	fmt.Printf("My new city is %s\n", person.city)
}
