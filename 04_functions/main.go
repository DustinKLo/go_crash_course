package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println(greeting("Dustin Lo you fucking suck"))
	fmt.Println(getSum(3, 4))
}
