// // An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
// // A slice is formed by specifying two indices, a low and high bound, separated by a colon: a[low : high]
// This selects a half-open range which includes the first element, but excludes the last one.

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	//	var s []int = primes[1:4] //slices
	//var s []int = primes[1:5]
	//var s []int = primes[0:1] //2
	//var s []int = primes[0:0] //[]
	var s []int = primes[:] //[2 3 5 7 11 13]
	fmt.Println(s)
}
