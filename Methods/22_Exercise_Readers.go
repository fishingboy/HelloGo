package main

import "golang.org/x/tour/reader"

type MyReader struct {
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
// @todo: 看不懂，之後再回頭來做
func main() {
	reader.Validate(MyReader{})
}
