package main

import (
	"fmt"
	"strconv"
)

// define Person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and i am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method, alters "objects" data (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		fmt.Println(p.firstName + "got married and her last name is now: " + p.lastName + " -> " + spouseLastName)
		p.lastName = spouseLastName
	}
}

func main() {
	// init person struct
	person1 := Person{
		firstName: "Dustin",
		lastName:  "Lo",
		city:      "Los Angeles",
		gender:    "m",
		age:       28,
	}
	fmt.Println(person1)
	fmt.Println(person1.age)

	//
	person3 := Person{"Bob", "Smith", "Los Angeles", "m", 29}
	fmt.Println(person3)

	// can also create structs like this
	person2 := Person{"Samantha", "William", "Los Angeles", "f", 28}
	fmt.Println(person2)
	person2.age++
	fmt.Println(person2)
	fmt.Println(person2.greet())
	person2.hasBirthday()
	fmt.Println(person2.greet())

	person2.getMarried("Smith")
	fmt.Println(person2.greet())

}
