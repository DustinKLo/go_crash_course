package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	var msg = "hello"

	go func() {
		wg.Add(1)
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "goodbye"
	time.Sleep(100 * time.Millisecond)
	wg.Wait()
}
