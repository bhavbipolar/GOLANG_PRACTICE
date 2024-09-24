//note --- in golang function argument associativity is from left to right whereas in C,C++ there is no associativity of function arument like (printf)

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		//if v := math.Pow(x, n) is dinitialization and declaration way
		//v<lim is putting condition
		return v
	} else {
		fmt.Printf("%g > =%g\n", v, lim) //%g is like %d
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

//for pow(3,2,10) under if loop result is x=3,n =2, lim =10, 3^2 =9 ,9<10 true so return 9
//for pow(3,3,20) under if loop result is x=3,n=3,lim =20, 3^3 =27, 27 <20 false so goes to else prints 27 >=20
//comes out of loop prints return lim ,i.e 20 then prints 9 as it was stored earlier in memory
