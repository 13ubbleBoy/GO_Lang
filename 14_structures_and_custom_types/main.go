package main

import "fmt"

func main() {
	mybill := newbill("praveen")
	fmt.Println("Bill:", mybill, "\n")
	fmt.Printf("Type of bill: %T\n\n", mybill)

	// accessing a single element
	fmt.Println(mybill.name, "\n")
}
