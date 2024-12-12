/*

--> panic is a tool for handling unexpected, critical errors in Go.
--> It unwinds the stack, executes deferred functions, and stops the program unless recovered.
--> Use panic sparingly and prefer returning errors for normal error handling.

--> panic
--> When panic is called, the program starts unwinding the stack, executing all deferred function calls.
--> It is better suited for scenarios where the program cannot proceed (e.g., invalid configuration,
	critical failures).

--> differ
--> All defer statements in the current goroutine are executed before the program crashes.
	Or after panic is executed.
--> Deferred functions are executed in reverse order when a panic occurs.

--> recovery
--> The recover function catches the panic and prevents the program from crashing.
--> After recovery, execution resumes normally from the point after the deferred recover function.

--> uses of recovery
	. Situations where the program cannot continue, such as corrupted state or critical missing resources.
	. To highlight unexpected errors during testing.
	. Libraries may use panic internally for serious issues but should document it properly.

*/

package main

import "fmt"

// program execution will continue
func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in safeFunction:", r)
		}
	}()
	fmt.Println("Executing safeFunction\n")
	panic("Panic in safeFunction\n")
	fmt.Println("This won't be printed")
}

func main() {
	// panic
	/*
		fmt.Println("Starting the program...")
		panic("Something went wrong!")
		fmt.Println("This won't be printed.") // this will not get executed because program will crash before this line.
	*/

	// panic with differ
	/*
		defer fmt.Println("Deferred function 1")
		defer fmt.Println("Deferred function 2")

		fmt.Println("Before panic")
		panic("Critical error!")
		fmt.Println("This will not be printed")
	*/

	// Recovering from Panic using recover function
	fmt.Println("\nStarting the program...\n")

	safeFunction()

	fmt.Println("Program continues after safeFunction\n")

}
