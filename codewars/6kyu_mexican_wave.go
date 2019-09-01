package main

import (
	"strings"
)

// Wave ....
func Wave(words string) (rc []string) {
	rc = []string{}
	letters := strings.Split(words, "")

	for i, l := range letters {
		if l != " " {
			tmp := make([]string, len(letters))
			copy(tmp, letters)
			tmp[i] = strings.Title(l)
			rc = append(rc, strings.Join(tmp, ""))
		}
	}

	return
}
