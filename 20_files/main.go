/*
--> bufio is package used for accepting user input.
	. bufio.NewReader() is used to read information, but have to specify from where it will read.
	. we read input from terminal for standard inputs. So we need to to use another package
	  called os and its function named "os.stdin".
*/

package main

import (
	"bufio"
	"fmt"
	"os" // to speficy terminal
	"strconv"
	"strings"
)

// function different type of inputs
func getInput(prompt string, read *bufio.Reader) (string, error) { // "(string, error)" return type of the function
	fmt.Print(prompt)

	input, err := read.ReadString('\n') // whenever there is a new line input it will stop reading.

	return strings.TrimSpace(input), err // "strings.TrimSpace(input)" it removes white spaces jo ki input k aage or
	// peeche ho wahi remove krta hai, input k between koi space ho toh wo remove nhi krega.
	// Ex: ->   prav  een
	// starting k 2 space ye remove kr dega but between v and e k nhi krega.
}

// getting the user input
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Enter your name: ")
	// name, _ := reader.ReadString('\n') // _ underscore tab use karte hai jab hume value nhi chaihiye hoti hai.
	// name = strings.TrimSpace(name)

	name, _ := getInput("Enter your name: ", reader) // getting name from the user using a functon

	b := newbill(name)
	fmt.Println("\nBill created for:", b.name, "\n")

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose your option\n1. Add a food item\n2. Tip us\n3. Get bill\nEnter your choice: ", reader)

	switch option {
	case "1":
		name, _ := getInput("Item name: ", reader)
		priceString, _ := getInput("Item price: ", reader) // input are always in string and we cannot add a string to a float32

		price, error := strconv.ParseFloat(priceString, 32) // 32 means it will get converted into float32
		// error will be if the method cannot parse a value. And values it cannot parse are like characters, strings etc.

		if error != nil {
			fmt.Println("The price must be a number!\n")
			promptOptions(b)
		}
		b.addItem(name, float32(price))

		fmt.Println("Item added: ", name, price, "\n")

		promptOptions(b)
	case "2":
		tipString, _ := getInput("Enter tip amount (₹): ", reader)

		tip, error := strconv.ParseFloat(tipString, 32) // string conversion into float

		if error != nil {
			fmt.Println("The tip must be a number!\n")
			promptOptions(b)
		}
		b.addTip(float32(tip))

		fmt.Println("Tip added:", tip, "\n")

		promptOptions(b)
	case "3":
		b.generateBill()
	default:
		fmt.Println("Invalid input, please try again\n")
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)

}
