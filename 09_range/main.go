package main

import "fmt"

func main() {
	fmt.Println("")
	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 33, 44, 55, 66, 77, 88, 99, 100}

	// loop through ids
	fmt.Println("using range")
	for i, id := range ids {
		fmt.Println(i, id)
	}

	// not using range
	fmt.Println("\nold school way")
	for i := 0; i <= len(ids)-1; i++ {
		fmt.Println(i, ids[i])
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	// range with map
	fmt.Println("\nrange with maps")
	emails := map[string]string{
		"bob":    "bob@gmail.com",
		"sharon": "sharon@gmail.com",
		"dustin": "dustin@gmail.com",
	}

	for k, v := range emails {
		fmt.Println(k, v)
	}
}
