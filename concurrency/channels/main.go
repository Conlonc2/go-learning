// small app to demo how chan's can be used to communicate data across threads (go rotuines)
/*
Overall what this app does:
	1. Create a channel for communication between threads
	2. We then call double(), this sleeps 2 seconds and then returns double the value passed into the function and
	   relays it back to a varible that is waiting on it
	3. While waiting for the return value (This is blocking) we run another go rotutne that will increment count while waiting on channel
	   a. The waiting function will be terminated once the main thread ends, we only use this to see that channel return values are blocking the termination of main thread
*/

package main

import (
	"fmt"
	"time"
)

func double(c chan int, val int) {
	time.Sleep(time.Second * 2)
	c <- 2 * val
}

func waiting() {
	count := 0
	for {
		time.Sleep(time.Millisecond * 30)
		fmt.Printf("Count: %d\n", count)
		count++
	}
}

func main() {
	channel := make(chan int)
	fmt.Println("Chan Val: ", channel)
	fmt.Println("Application Start...")
	go double(channel, 4)
	go waiting()
	retVal := <-channel

	fmt.Printf("%d\n", retVal)
	fmt.Println("Application End...")

}
