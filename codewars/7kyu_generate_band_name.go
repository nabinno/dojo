package main

import (
	"fmt"
	"strings"
)

// GenerateBandName ...
func GenerateBandName(s string) (rc string) {
	if s[0] == s[len(s)-1] {
		rc = fmt.Sprintf("%v%v", strings.Title(s), s[1:len(s)])
	} else {
		rc = fmt.Sprintf("The %v", strings.Title(s))
	}
	return
}
