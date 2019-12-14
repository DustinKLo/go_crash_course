package main

import (
	"fmt"
	"reflect"
)

/*
A VALUE CAN HAVE MORE THAN ONE TYPE
A VALUE CAN HAVE MORE THAN ONE TYPE
A VALUE CAN HAVE MORE THAN ONE TYPE
A VALUE CAN HAVE MORE THAN ONE TYPE
A VALUE CAN HAVE MORE THAN ONE TYPE
*/

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

type human interface {
	speak()
	Name() string
}

func (p person) Name() string {
	return fmt.Sprintf("function foo: %s %s", p.first, p.last)
}

// how to access struct field from interface
// https://stackoverflow.com/a/52664060
func foo(h human) {
	fmt.Println(h.Name())
	// fmt.Println("i called human", reflect.TypeOf(h))
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into barrrrrr", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into barrrrrr", h.(secretAgent).first)
	}
	fmt.Println("I was passed into bar", h)
}

func nativeSpeak(h human) {
	h.speak()
}

type hotdog int

func main() {
	sa1 := secretAgent{
		person: person{"james", "bond"},
		ltk:    true,
	}

	sa2 := secretAgent{
		person: person{"miss", "moneypenny"},
		ltk:    true,
	}

	p1 := person{
		first: "dr.",
		last:  "yes",
	}

	fmt.Println(sa1)
	fmt.Println(p1)

	fmt.Println("########################################")
	fmt.Println("netive speak method")
	sa1.speak()
	sa2.speak()
	p1.speak()

	fmt.Println("\n########################################")
	fmt.Println("interface speak function")
	bar(sa1)
	bar(sa2)
	bar(p1)
	fmt.Println("########################################\n")

	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println("\n############################")
	foo(sa1)
	foo(p1)
	nativeSpeak(sa1)
	nativeSpeak(p1)

	fmt.Println(reflect.TypeOf(sa1))
	fmt.Println(reflect.TypeOf(p1))
	fmt.Println("############################\n")
}
