package main

import "fmt"

func main() {
	mybill := newbill("praveen")

	mybill.addItem("tea", 20)
	mybill.addItem("bread pakoda", 30)

	mybill.addTip(30)

	fmt.Println(mybill.format())
	fmt.Println()

}
