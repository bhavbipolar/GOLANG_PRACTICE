// If the slice is backed by the array and arrays are fixed length, then how is that possible a slice is a dynamic length?.
// Well, the answer is when the new elements append to the slice a new array is created. Now, the elements present in the existing array are copied into a new array and return a new slice reference to this array(new array).

package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	//apend work on nil slices
	s = append(s, 0)
	printSlice(s)

	//the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	//we can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
