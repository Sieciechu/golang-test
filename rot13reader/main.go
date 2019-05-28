package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)

	for i, c := range p {
		if 0 == c {
			continue
		}
		if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
			p[i] += 13
		}
		if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
			p[i] -= 13
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
