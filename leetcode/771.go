package main

import "strings"

// NumJewelsInStones (771)
func NumJewelsInStones(J string, S string) int {
	var gems []string

	for i := 0; i < len(S); i++ {
		if strings.Contains(J, string(S[i])) {
			gems = append(gems, string(S[i]))
		}
	}
	return len(gems)
}
