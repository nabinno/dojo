package main

import "golang.org/x/tour/reader"

type myReader struct{}

func (r myReader) Read(b []byte) (n int, err error) {
	for n, err = 0, nil; n < len(b); n++ {
		b[n] = 'A'
	}
	return
}

// Validate (methods 22)
func Validate() {
	reader.Validate(myReader{})
}
