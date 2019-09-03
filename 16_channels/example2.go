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
	ch := make(chan int, 50) // creating a buffer channel (allowing 50 spare)

	wg.Add(2)

	go func(ch <-chan int) { // specifying list only channel
		// for i := range ch { // one way to loop through channel
		// 	fmt.Println(i)
		// }
		for { // another way to loop through channel
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // send only channel
		ch <- 42
		ch <- 27
		close(ch) // close channel, unable to push new messages
		wg.Done()
	}(ch)

	wg.Wait()
}
