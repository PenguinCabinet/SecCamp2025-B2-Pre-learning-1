package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (self MyReader) Read(v []byte) (int, error) {
	v[0] = byte('A')
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
