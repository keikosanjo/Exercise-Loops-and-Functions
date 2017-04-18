package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g ", e)

}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	Z := z
	for i := 0; i < 10; i++ {

		z = z - (math.Pow(z, 2)-x)/(2*z)

		if math.Abs(Z-z) <= 1e-10 {
			break
		}
		Z = z
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
}
