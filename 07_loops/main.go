package main

import "fmt"

func main() {

	// while loop
	i := 0
	for i < 5 {
		fmt.Println("Value of x:", i)
		i++
	}
	fmt.Println()

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Value of x:", i+1)
	}
	fmt.Println()

	// loop though slice
	names := []string{"Praveen", "Ankit", "Piyush", "Jeetu", "Mithun"}
	for i := 0; i < 5; i++ {
		fmt.Println(names[i])
	}
	fmt.Println()

	// for each loop
	for index, value := range names {
		fmt.Printf("Index %v and value %v\n", index, value)
	}
	fmt.Println()

	// what if in the above loop we don't need index --> we replace that parameter with an underscore "_"
	for _, value := range names {
		fmt.Printf("Value: %v\n", value)
	}
	fmt.Println()

	// what happen if we update the 'value' parameter in the above, will it affect the original array
	for _, value := range names {
		fmt.Printf("Value: %v\n", value)
		value = "new value" // this does not affect the original slice, value is a local copy.
	}
	fmt.Println(names)
	fmt.Println()
}
