package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g ", e)

}

func main() {
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
}
