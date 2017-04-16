package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for j := 0; j < dy; j++ {
		p[j] = make([]uint8, dx)
	}

	return p
}

func main() {
	pic.Show(Pic)
}
