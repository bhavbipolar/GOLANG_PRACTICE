package main

import "fmt"

// newCounter function to
// isolate global variable
func newCounter() func(z int) int {
	GFG := 0
	return func(z int) int {
		GFG += z
		return GFG
	}
}
func main() {
	// newCounter function is
	// assigned to a variable
	counter, counter1 := newCounter(), newCounter()
	i := 1
	e := counter(i)
	j := counter(i)
	k := counter(i)
	//counter1 := newCounter()

	l := counter1(-2 * i)   // counter1(-2) ,l=-2 , GFG = -2
	m := counter1(-2*i + 4) // counter1(2),m = 0, gfh = 0
	n := counter1(-2*i - 4) //counter1(-6),n = -6, gfg = -6

	fmt.Println(e, j, k, l, m, n)

}
