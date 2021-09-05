package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(rb []byte) (n int, e error) {
	n, e = reader.r.Read(rb)
	for i := range rb {
		rb[i] = lot13(rb[i])
	}
	return
}

func lot13(b byte) byte {
	if ('A' <= b && b <= 'L') || ('a' <= b && b <= 'l') {
		return b + 13
	} else if ('A' <= b) && (b <= 'z') {
		return b - 13
	}
	return b
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
