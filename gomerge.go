package gomerge

import (
	"fmt"
)

// CustomPrinter struct that holds additional functionality
type CustomPrinter struct{}

// Println method that wraps fmt.Println
func (cp CustomPrinter) Println(a ...interface{}) (n int, err error) {
	// Call the original Println function
	n, err = fmt.Println(a...)

	// Perform additional actions after Println completes
	cp.afterPrintln()

	return n, err
}

// afterPrintln performs additional actions after Println completes
func (cp CustomPrinter) afterPrintln() {
	fmt.Println("This message is printed after the original Println.")
}
