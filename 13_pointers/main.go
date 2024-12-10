/*
-->
*/

package main

import "fmt"

func updateName(ptr *string, name string) {
	*ptr = name
}

func main() {
	name := "Praveen"
	fmt.Println("Memory location of the variable name:", &name, "\n")

	ptr := &name
	/*
		--> 'ptr' has its own memory address and in that address it stores the address of vaiable
		name as a value. It means we can say that 'pt' has a value which is the address of name.
	*/
	fmt.Printf("Value at ptr: %v\nAddress of pts: %v\n", ptr, &ptr)
	fmt.Printf("Value in the address that ptr holds: %v\n\n", *ptr)

	// now if we change the value at address stored by ptr, the value inside name changes
	*ptr = "Piyush"
	fmt.Println(name, "\n")

	// passing pointer to a funciton
	updateName(ptr, "Mithun")
	fmt.Println(name, "\n")

	// or we can directly pass the address
	updateName(&name, "Jeetu")
	fmt.Println(name, "\n")

}
