package main

import (
	"strings"
)

func main() {
}

// DefangIPaddr (1108)
func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

// NumJewelsInStones (771)
// @example
//    ```
//    j := "aA"
//    s := "aAAbbbb"
//    rc := NumJewelsInStones(j, s)
//    fmt.Println(rc)
//    ```
func NumJewelsInStones(J string, S string) int {
	var gems []string

	for i := 0; i < len(S); i++ {
		if strings.Contains(J, string(S[i])) {
			gems = append(gems, string(S[i]))
		}
	}
	return len(gems)
}
