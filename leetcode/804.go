package main

import (
	"bytes"
	"fmt"
)

// UniqueMorseRepresentations ....
func UniqueMorseRepresentations(words []string) int {
	morseCodes := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	rcMap := make(map[string]bool, len(words))
	for _, w := range words {
		var b bytes.Buffer
		for i := 0; i < len(w); i++ {
			fmt.Fprint(&b, morseCodes[w[i]-'a'])
		}
		rcMap[b.String()] = true
	}
	return len(rcMap)
}
