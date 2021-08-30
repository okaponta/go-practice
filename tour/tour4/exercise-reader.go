package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (a MyReader) Read(rb []byte) (n int, e error) {
	for i := range rb {
		rb[i] = 'A'
	}
	return len(rb), nil
}

func main() {
	reader.Validate(MyReader{})
}
