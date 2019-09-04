package main

import (
	"fmt"
)

func main() {
	fmt.Println([]byte("hello go!"))

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface {
	Write([]byte) (int, error) // method takes []byte, returns int and error
}

type ConsoleWriter struct{}

// method to ConsoleWriter struct
// argument is []byte and returns (int, err)
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
