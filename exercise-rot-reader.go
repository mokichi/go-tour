package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	n, err := rot13.r.Read(b)
	for i, c := range b {
		switch {
		case 'A' <= c && c <= 'M', 'a' <= c && c <= 'm':
			b[i] += 13
		case 'N' <= c && c <= 'M', 'n' <= c && c <= 'z':
			b[i] -= 13
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
