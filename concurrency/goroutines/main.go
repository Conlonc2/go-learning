// Simple application to demonstrate wait groups so that go routines are synced and given enough time to run
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(something string) {
	wg.Add(1)
	fmt.Println(something)
	wg.Done()
}

func main() {
	go say("Hello")
	say("How")
	go say("are")
	say("you")
	wg.Wait()
}
