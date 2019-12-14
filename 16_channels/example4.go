package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time time.Time

	severity string
	message  string
}

var wg = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)

func main() {
	wg.Add(2)
	go addToLogger()
	go logger()
	wg.Wait()
}

func addToLogger() {
	logCh <- logEntry{time.Now(), logInfo, "app is starting"}
	logCh <- logEntry{time.Now(), logInfo, "app is shutting down"}
	close(logCh)
	wg.Done()
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
	}
	wg.Done()
}
