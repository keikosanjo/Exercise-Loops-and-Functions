package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := float64(1)
	var I int
	for i := 0; i < 1000; i++ {
		var Z float64

		Z = z
		z = z - (math.Pow(z, 2)-x)/(2*z)

		if math.Abs(Z-z) <= 1e-10 {
			I = i
			break
		}

	}
	return z, I

}

func main() {
	z, i := Sqrt(2)
	fmt.Println(i)
	fmt.Println(z)
	fmt.Println(math.Sqrt(2))
}

