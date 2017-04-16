package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64) {
	z := float64(1)
	for i := 0; i < 1000; i++ {

		Z := z
		z = z - (math.Pow(z, 2)-x)/(2*z)

		if math.Abs(Z-z) <= 1e-10 {
			break
		}

	}
	return z

}

func main() {
	z := Sqrt(2)
	fmt.Println(z)
	fmt.Println(math.Sqrt(2))
}

