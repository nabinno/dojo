package main

import (
	"fmt"
	"strings"
)

func main() {
	j := "aA"
	s := "aAAbbbb"
	rc := numJewelsInStones(j, s)
	fmt.Println(rc)
}

// Defanging an IP Address (1108)
func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// Jewels and Stones (771)
func numJewelsInStones(J string, S string) int {
	var gems []string

	for i := 0; i < len(S); i++ {
		if strings.Contains(J, string(S[i])) {
			gems = append(gems, string(S[i]))
		}
	}
	return len(gems)
}
