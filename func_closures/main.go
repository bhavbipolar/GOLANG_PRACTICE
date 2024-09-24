package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		// //fmt.Println(
		// 	pos(i),
		// 	neg(-2*i),
		// )
		v1, v2 := pos(i), neg(-2*i)
		fmt.Println(v1, v2)

	}
}

/*
loop 1 : i=0,
			pos(0),neg(0)
			sum =0, sum =0
			v1 =0, v2=0
loop 2 : i=1,
			pos(1),neg(-2)
			sum =1, sum =-2
			v1 =1, v2=-2
loop 3 : i=2,
			pos(2),neg(-4)
			sum =3, sum =-6
			v1 =3, v2=-6
loop 4 : i=3,
			pos(3),neg(-6)
			sum =6, sum =-12
			v1 =6, v2=-12

loop 5 : i=4,
			pos(4),neg(-8)
			sum =10, sum =-20
			v1 =10, v2=-20
*/
