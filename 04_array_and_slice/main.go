package main

import "fmt"

func main() {
	// arrays: fixed length and can hold same type of data like int, float etc.
	var nums [3]int = [3]int{1, 5, 3}
	fmt.Println(nums, "\n")

	// 2nd way and length of the array
	var nums2 = [3]int{4, 5, 3}
	fmt.Println("Array:", nums2, "\nLength:", len(nums2), "\n") // finding length of the array

	// 3rd way
	names := [4]string{"Praveen", "Ankit", "Piyush", "Jeetu"}
	fmt.Println("Array:", names)
	fmt.Printf("Type of names: %T\n\n", names)

	// Slices: Variable length array and also modifyable
	var scores = []int{21, 32, 24, 67} // as the square brackets is empty, it automatically sets it to be an slice
	fmt.Println(scores)
	fmt.Printf("Type of scores: %T\n\n", scores)

	// edit slice
	scores[1] = 30
	fmt.Println(scores, "\n")

	// add an element: 'append()' this function does not modify the existing slice, instead it return a new slice.
	scores = append(scores, 56) // we over writing the scores, as append return a new slice.
	fmt.Println(scores)
	fmt.Println("Slice length:", len(scores), "\n")

	// Slice ranges: It is a way to get a range of elements from an existing array.
	range1 := names[0:2] // 0 to 2-1 elements are included
	fmt.Println(range1, "\n")

	range2 := scores[1:] // this means from index 1 to all
	fmt.Println(range2)
	fmt.Println("Length of range:", len(range2), "\n")

	range3 := scores[:4] // this means from index 0 to 3
	fmt.Println(range3, "\n")

	range1 = append(range1, "Mithun")
	fmt.Println(range1, "\n")

}
