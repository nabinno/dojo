package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (n int, err error) {
	n, err = rot.r.Read(b)
	for i := 0; i < len(b); i++ {
		char := b[i]
		if (char >= 'A' && char < 'N') || (char >= 'a' && char < 'n') {
			b[i] += 13
		} else if (char > 'N' && char <= 'Z') || (char > 'm' && char <= 'z') {
			b[i] -= 13
		}
	}
	return
}

// Rot13 (methods 23)
func Rot13(s string) {
	sr := strings.NewReader(s)
	rot := rot13Reader{sr}
	_, _ = io.Copy(os.Stdout, &rot)
}
