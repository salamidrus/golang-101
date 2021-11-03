package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (valuen receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// // Init person using struct
	person1 := Person{firstName: "Idrus", lastName: "Salam", city: "Jogja", gender: "f", age: 26}
	// // ALternative
	person2 := Person{"Bob", "Johnson", "Jogja", "m", 50}

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)
	person1.hasBirthday()
	person1.getMarried("Putri")

	person2.getMarried("Alisyah")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
