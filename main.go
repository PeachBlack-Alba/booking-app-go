package main

import "fmt"

// We need to define the endpoint so Go knows where to start executing the code
// If we have multiple files in the project, we need to give Go a compiler starting point
// Meaning the first line of code where the execution starts
// The main function is the entrypoint of a Go program:
// For 1 Go application we have 1 main func

func main() {
	// Print is a method from the fmt package, so we need to call it like so:
	fmt.Print("Hello world")
}
