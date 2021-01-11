package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}

	w.Write([]byte("print me out, please"))

	// --- empty interface
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is an unknown type")
	}
}

// --- interface declaration syntax
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// --- embed smaller interfaces into bigger one
type WriterCloser interface {
	Writer
	Closer
}

// --- naming convention
// if the interface has one method the name should be functionName+er e.g. Write+er = Writer

type ConsoleWriter struct{}

// --- interface implementation has to implement functions of the interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// --- implementing with values vs pointers
// methods set of value is all methods with value receivers
// methods set of pointers is all methods regardless of receiver type

// --- BEST PRACTICES
// use many small interfaces
// dont export interfaces for types that will be consumed
// export interfaces for types that will be used by package
// design funcions and methods to receive interfaces whenever possible
