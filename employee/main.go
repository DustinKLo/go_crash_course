package main

import "fmt"

type Employee struct {
	name    string
	salary  float64
	contact Contact
}

type Contact struct {
	phone string
}

func (e *Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) raiseSalary(percentage float64) float64 {
	if percentage < 1 {
		// fmt.Println("You're cutting their pay? Not cool.")
		percentage++
	}
	e.salary = e.salary * percentage
	return e.salary
}

func (e *Employee) changePhone(newPhone string) {
	e.contact.phone = newPhone
}

func main() {
	e := Employee{
		name:    "Ross Geller",
		salary:  1200,
		contact: Contact{phone: "16263211938"},
	}

	fmt.Println(e)
	e.changeName("Monica Geller")
	e.raiseSalary(0.99)
	fmt.Println(e)
	e.raiseSalary(1.5)
	e.changePhone("123456789")
	fmt.Println(e)
}
