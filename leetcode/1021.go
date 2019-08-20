package main

import "bytes"

// RemoveOuterParentheses ...
func RemoveOuterParentheses(S string) string {
	var res bytes.Buffer
	cnt := 0
	for _, c := range S {
		if cnt != 0 && !(cnt == 1 && c == ')') {
			res.WriteRune(c)
		}
		if c == '(' {
			cnt++
		} else {
			cnt--
		}

	}
	return res.String()
}
