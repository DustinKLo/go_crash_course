package main

import (
	"fmt"
	"sync"
)

/*
channel basics
restricing data flow
buffered channels
closing channels
for...range loops with channels
select statements
*/

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // make channel (strongly typed)

	wg.Add(2)

	go func(ch <-chan int) { // specifying list only channel
		i := <-ch // receiving data from the channel
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // send only channel
		x := 42
		ch <- x // passing data into the channel
		// fmt.Printf("data passed into channel: %d\n", x)
		wg.Done()
	}(ch)

	wg.Wait()
}
