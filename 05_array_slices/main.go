package main

import "fmt"

func main() {
	// // arrays
	// var fruitArr [2]string

	// // assign values
	// fruitArr[0] = "apple"
	// fruitArr[1] = "orange"

	// // declare and assign
	// fruitArr := [2]string{"apple", "orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitArr := []string{"apple", "orange", "grape"}
	fmt.Println(fruitArr)
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[1:2])
}
