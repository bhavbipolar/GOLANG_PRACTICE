// The range keyword is mainly used in for loops in order to iterate over all the elements of a map, slice, channel, or an array. When it iterates over the elements of an array and slices then it returns the index of the element in an integer form. And when it iterates over the elements of a map then it returns the key of the subsequent key-value pair. Moreover, range can either returns one value or two values

package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //==2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
