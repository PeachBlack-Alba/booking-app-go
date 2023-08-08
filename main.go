package main

import "fmt"

// We need to define the endpoint so Go knows where to start executing the code
// If we have multiple files in the project, we need to give Go a compiler starting point
// Meaning the first line of code where the execution starts
// The main function is the entrypoint of a Go program:
// For 1 Go application we have 1 main func

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	// Print is a method from the fmt package, so we need to call it like so:
	fmt.Println("Welcome to", conferenceName, " our conference booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
	// Print formatted data: fmt.printft()
	// It takes a template string that contains the text that needs to be formatted
	// Plus some annotation verbs (or placeholder) that tells the fmt functions
	// How to format the variable passed in
	fmt.Printf("Welcome to %v our conference booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	var userName string
	var userTickets int

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}

/*Basic Data Types
Strings:
	- For textual data, defined with double quotes, eg: "This is a string"
Integers
	- Representing whole numbers, positive and negative, eg: 5, 120, -20
	- There are many more numeric data types
*/

/*Go is a statically typed language
- You need to tell Go Complier, the data type when declaring the variable
- Type Inference: BUT, Go can infer the type when you assing a value
*/

/*Compared to Java:
Go	     Java
int8 => byte
int16 => short
int32 => int
int64 => long
*/
