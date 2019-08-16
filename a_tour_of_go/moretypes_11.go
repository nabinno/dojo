package main

import "fmt"

// PrintSlice (moretypes 11)
func PrintSlice(s []int) string {
	return fmt.Sprintf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
