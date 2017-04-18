package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if err != nil {
	} else {
		for i := 0; i < n; i++ {
			if b[i] >= 'a' && b[i] <= 'm' {
				b[i] = b[i] + 13
			} else if b[i] > 'm' && b[i] <= 'z' {
				b[i] = b[i] - 13
			} else if b[i] >= 'A' && b[i] <= 'M' {
				b[i] = b[i] + 13
			} else if b[i] > 'M' && b[i] <= 'Z' {
				b[i] = b[i] - 13

			}
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
