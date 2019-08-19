package main

import "strings"

// DefangIPaddr (1108)
func DefangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}
