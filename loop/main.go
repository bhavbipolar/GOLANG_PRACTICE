// //z -= (z*z - x) / (2*z)
// Implement this in the func Sqrt provided. A decent starting guess for z is 1, no matter what the input. To begin with, repeat the calculation 10 times and print each z along the way. See how close you get to the answer for various values of x (1, 2, 3, ...) and how quickly the guess improves.
//newton -Raphson method to find closes approximation to sqrt of a number

package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	iter := 20
	for i := 0; i < iter; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
