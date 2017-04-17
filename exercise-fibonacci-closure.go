package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	s := 0
	m := 1
	sum := 0
	return func(x int) int {
	if x == 0 {
			return 0
		} else if x == 1 {
			return 1
		}
		sum = s + m
		s = m
		m = sum
		return sum
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
