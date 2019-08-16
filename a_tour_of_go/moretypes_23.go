package main

import "strings"

// WordCount (moretypes 23)
func WordCount(s string) map[string]int {
	rc := make(map[string]int)
	for _, w := range strings.Fields(s) {
		rc[w]++
	}
	return rc
}
