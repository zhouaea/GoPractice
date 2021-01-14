package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(b [] byte) (int, error) {
	for k, _ := range b {
		b[k] = 'A'
	}
	
	return len(b), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
