/*

--> No, Go (Golang) is not a single-threaded language. It is designed for concurrent programming
	and supports multithreading through a concept called goroutines.

--> Concurrency 							vs 			Parallelism
	Structuring a program to do multiple				Running multiple tasks at the
	tasks independently (not necessarily				same time, typically on multiple CPUs.
	simultaneously).

--> Concurrency in Go is a fundamental concept, and the language provides robust support for it
	through goroutines and channels. Go’s concurrency model is lightweight and designed to enable
	programs to handle multiple tasks simultaneously.



--> Components <--

--> Goroutines:
	. Lightweight threads managed by the Go runtime.
	. Created using the go keyword.
	. Thousands of goroutines can run concurrently in a single program.

--> Channels:
	. Used for communication and synchronization between goroutines.
	. Provide a way to safely share data.

--> Select Statement:
	. Waits for multiple channel operations, enabling responsive and concurrent code.

--> sync Package:
	. Provides primitives like sync.Mutex for mutual exclusion.


*/

package main

import (
	"fmt"
	"time"
)

func printNumbers(i int, max int) {
	fmt.Println()
	for ; i <= max; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}

func main() {
	go printNumbers(1, 10) // a goroutine is created by prefixing a function call with the go keyword.
	fmt.Println("I am the 1st fmt.Println() from main function\n")
	go printNumbers(1, 20)

	time.Sleep(2 * time.Second)
	fmt.Println("I was printed after 2 seconds\n") // ye 2 second baad print hoga

	go printNumbers(1, 30) // l1
	go printNumbers(1, 40) // l2
}

/*

--> Why l1 and l2 are not getting executed
--> Because the main function is terminating before the goroutines have a chance to run.
	In Go, when the main function of your program exits, the entire program terminates,
	including any running or scheduled goroutines.

	. The time.Sleep(2 * time.Second) gives only 2 seconds for the previous goroutines to execute.
	. After the sleep, the main function ends without waiting for the new goroutines
	  (go printNumbers(1, 30) and go printNumbers(1, 40)) to complete.
	. Goroutines run concurrently, but their execution depends on the Go scheduler. If the main
	  function exits before the goroutines get a chance to execute, they won't run.

--> To solve it we have to sync them.

*/
