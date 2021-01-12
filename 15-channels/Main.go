package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// --- syntax declaring a channel
// channels are strongly typed
//ch := make(chan int)

// by default channel is unbuffored
// how to solve it?
var ch = make(chan int, 50)
var signal = make(chan struct{}) // empty channel with no data fields
// used only for signaling

func main() {

	// --- create waiting group of two
	wg.Add(2)

	// 1st go routine in the waiting group
	// --- passing channels with direction restriction (read)
	go func(ch <-chan int) {
		// read data from the channel into variable "i"
		i := <-ch
		fmt.Println(i)
		i = <-ch
		fmt.Println(i)
		// /ch <- 18 !!! will not work
		wg.Done()
	}(ch)

	// 2nd go routine in the waiting group
	// --- passing channels with direction restriction (write)
	go func(ch chan<- int) {
		// write data to the channel
		ch <- 18
		ch <- 36
		// two writes to channel, we need a buffer
		//close(ch)
		//i := <-ch !!! will not work
		wg.Done()
	}(ch)

	// --- how i understand it
	// we have two go routines
	// they can execute in a random order
	// if the 1st reads from an empty channel, it stops until the channel is filled with data
	// so 2nd go routine can execute before or after the 1st
	// because the 1st will wait

	wg.Wait()

	go looper()
	ch <- 11
	ch <- 123
	signal <- struct{}{}

}

// --- closing a channel
// close(ch)
// can be done only once, channel cannot be reopened

func looper() {
	for {
		// --- select statement
		// block if all channels blocked
		select {
		case data := <-ch:
			fmt.Printf("looper printed: %v\n", data)
		case <-signal:
			break
		}
	}
}
