package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(c chan int, value int) {
	defer wg.Done()
	c <- value // * 5
}

func main() {

	len := 1000
	fooChannel := make(chan int, len) // buffer

	for i := 0; i < len; i++ {
		wg.Add(1)
		go foo(fooChannel, i)
	}

	wg.Wait()
	close(fooChannel)

	// go func() { // allows you to not need a buffer in channel
	// 	wg.Wait()
	// 	close(fooChannel)
	// }()

	for item := range fooChannel {
		fmt.Println(item)
	}

}
