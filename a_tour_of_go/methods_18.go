package main

import "fmt"

type ipaddr [4]byte

// fmt.Stringer (methods 18)
func (is ipaddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", is[0], is[1], is[2], is[3])
}
