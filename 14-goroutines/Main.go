package main

import (
	"fmt"
	"runtime"
	"sync"
)

// --- go routine is high level abstraction of a thread

// --- check for goroutine race at compile time
// go run -race Main.go

// --- wait group
var wg = sync.WaitGroup{}

// --- mutex - a lock
var m = sync.RWMutex{}

func main() {
	// --- listing threads
	// by default there is the same amount of them as CPU cores
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// synchronising go routine with the main
	wg.Add(1) //  increment number of go routines to wait

	// --- go routine
	go func() {
		fmt.Println("Print from the goroutine")
		wg.Done() // decrement number of go routines to wait
	}()

	// wait until the number of groups you are waiting for is zero
	wg.Wait()

	// --- mutex

	// the example with interface does not make sense
	// i just done it for practice
	var c writerIncrementer = &counter{0}

	for i := 0; i < 10; i++ {
		wg.Add(2)
		// lock outside the goroutine
		m.RLock()
		go c.write() // unlock inside a goroutine
		m.Lock()
		go c.increment()
	}
}

// interface
// ###########################
type incrementer interface {
	increment() int
}

type writer interface {
	write() (int, error)
}

type writerIncrementer interface {
	incrementer
	writer
}

// ###########################

// implementing interface on struct counter
// ##############################
type counter struct {
	val int
}

func (c *counter) increment() int {
	c.val++
	m.Unlock()
	wg.Done()
	return c.val
}

func (c counter) write() (int, error) {
	n, err := fmt.Println(c.val)
	m.RUnlock()
	wg.Done()
	return n, err
}

// ##############################

// --- dont create goroutines in libraries

// --- when creating goroutine - know how it will end
// you will avoid subtle memory leaks
