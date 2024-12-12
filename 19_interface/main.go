/*

--> An interface is a type that specifies a set of method signatures.
--> It allows you to define and work with types that share common behavior, enabling
	polymorphism and decoupling.
--> Interfaces are central to Go's design philosophy, emphasizing composition over inheritance.

--> The zero value of an interface is nil.
--> A nil interface means both the dynamic type and value are nil.

--> Basically 1 hi method ko hum different purpose k liye use kar sakte hai.

*/

package main

import (
	"fmt"
	"math"
)

// interface
type shape interface {
	BOUNDARY() float32
	AREA() float32
}

// structures of different shapes
type circle struct {
	radius float32
}

type rectangle struct {
	length  float32
	breadth float32
}

// function for area
func (c circle) AREA() float32 {
	area := math.Pi * c.radius * c.radius
	return area
}

// function for circumference
func (c circle) BOUNDARY() float32 {
	boundary := 2 * math.Pi * c.radius
	return boundary
}

// function for area
func (r rectangle) AREA() float32 {
	area := r.length * r.breadth
	return area
}

// function perimeter
func (r rectangle) BOUNDARY() float32 {
	boundary := 2 * (r.length + r.breadth)
	return boundary
}

// this print function will work in both the cases, for circle as well as rectangle
func printDetails(s shape) {
	fmt.Printf("\nLength of the Boundary of %T: %0.2f\n", s, s.BOUNDARY())
	fmt.Printf("Area of %T: %0.2f\n\n", s, s.AREA())
}

func main() {
	c := circle{
		radius: 15.3532,
	}

	r := rectangle{
		length:  43,
		breadth: 23,
	}

	printDetails(c)
	printDetails(r)
}
