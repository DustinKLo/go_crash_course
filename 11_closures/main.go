package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Println(sum, x, sum+x)
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i <= 10; i++ {
		fmt.Println(i, sum(i))
	}
}
