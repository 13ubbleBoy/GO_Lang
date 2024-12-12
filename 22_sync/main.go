/*

--> It is necesary to maintain synchronization between goroutines.

--> sync Package:
	. Provides primitives like sync.Mutex for mutual exclusion.

--> sync.WaitGroup:
	. A WaitGroup is used to wait for a collection of goroutines to finish.
	. The Add method increments the counter before starting a goroutine.
	. Each goroutine calls Done when it finishes, decrementing the counter.

--> Wait:
	. The Wait method blocks the main function until the counter becomes zero,
	  ensuring all goroutines complete their execution.

--> var wg sync.WaitGroup
	. This line declares a WaitGroup variable named wg.
	. A WaitGroup is a type provided by the sync package in Go to wait for a collection of
	  goroutines to finish executing.
	. The wg variable will be used to keep track of how many goroutines are currently running
	  and ensure that the main function or other goroutines wait until they all complete.

--> waitKro.Add(1)
	. This line increments the counter inside the WaitGroup by 1.
	. It indicates that one goroutine is about to start and must be accounted for in the WaitGroup.
	. You call wg.Add(n) to increment the counter by n, where n is the number of new goroutines
	  you plan to start.

--> How It Works Together:
	. Increment: When you plan to launch a new goroutine, call "waitKro.Add(1)" to increment the counter,
	  ensuring the main thread knows a new task is running.
	. Decrement: Inside the goroutine, you call "waitKro.Done()" when the task is complete. This decrements
	  the counter by 1.
	. Wait: Use "waitKro.Wait()" to block execution in the main function (or wherever it's called) until the
	  counter in wg reaches 0, meaning all goroutines have completed.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(i int, max int, waitKro *sync.WaitGroup) {
	for ; i <= max; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
	defer waitKro.Done()
}

func main() {
	var waitKro sync.WaitGroup

	waitKro.Add(1)
	go printNumbers(1, 10, &waitKro)
	fmt.Printf("\nI am the 1st fmt.Println() from main function\n")

	waitKro.Add(1)
	go printNumbers(1, 20, &waitKro)

	time.Sleep(2 * time.Second)
	fmt.Printf("\nI was printed after 2 seconds\n")

	waitKro.Add(1)
	go printNumbers(1, 30, &waitKro) // l1

	// eske liye alag se waitKro() use krna padega
	waitKro.Add(1)
	go printNumbers(1, 40, &waitKro) // l2

	waitKro.Wait()
	fmt.Printf("\nEverything executed!\n")
}

/*

--> Why we need to use waitKro.Add(1) before every go routine ?
	waitKro.Add(1) must be called before starting every goroutine because it ensures that the
	WaitGroup's counter is properly incremented to account for that goroutine.

	--> Calling Add(1) before starting the goroutine ensures the counter is incremented before the
		goroutine executes and calls Done().
	--> This avoids the race condition where Done() might decrement the counter before Add(1)
		increments it.

	--> If Done() is called on a WaitGroup before the corresponding Add(n) is executed, it leads to
		a panic because the counter becomes negative.
	--> If Done() is called when the counter is zero (or uninitialized for the goroutine), it triggers
		a panic:
	--> This happens because WaitGroup does not allow its counter to become negative.

	--> Add(1) signals the WaitGroup to track the new goroutine. Without this, the WaitGroup.Wait()
		might finish prematurely, not waiting for all goroutines to complete.

*/
