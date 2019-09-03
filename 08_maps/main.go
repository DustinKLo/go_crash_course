package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)

	fmt.Println(emails)
	// assign key value
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	// fmt.Println(emails["bob"])
	// fmt.Println(len(emails))

	// delete from map
	delete(emails, "bob")
	fmt.Println(emails)

	// declare map and add key value
	emails2 := map[string]string{
		"bob":    "bob@gmail.com",
		"sharon": "sharon@gmail.com",
	}
	fmt.Println(emails2)
}
