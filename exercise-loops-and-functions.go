package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var Z float64
	for {
		Z = z
		z = z - (math.Pow(z, 2)-x)/(2*z)

		if math.Abs(Z-z) < 1e-10 {
			return z
		}

	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

