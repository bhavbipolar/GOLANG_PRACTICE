package main

import "fmt"

func add(i, j int) int {
	return i + j
}

func main() {
	i, j := 2, 4
	j, k, r := i*4, j*-1, j+4 //j =8, k= -4,r=8
	s1 := add(i, j)           //add(2,8)
	s2 := add(i*2, j*-4)      //add(4,-32)
	s3 := add(j, k)           //add(8,-4)
	s4 := add(k, r)           //add(-4,8)

	fmt.Println(s1, s2, s3, s4)

}
