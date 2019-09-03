package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x < y {
		fmt.Printf("%d < %d\n", x, y)
	} else {
		fmt.Printf("%d >= %d\n", x, y)
	}

	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is neither red nor blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}
}
