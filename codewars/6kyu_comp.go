package main

import (
	"math"
	"reflect"
	"sort"
)

// Comp ....
func Comp(array1 []int, array2 []int) bool {
	if len(array1) == 0 || len(array2) == 0 {
		if len(array1) == len(array2) {
			return true
		}
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	for i := range array2 {
		f2 := float64(array2[i])
		chk := int(math.Sqrt(f2))
		if f2 == math.Pow(float64(chk), 2) {
			array2[i] = chk
		} else {
			return false
		}
	}

	for j := range array1 {
		array1[j] = int(math.Abs(float64(array1[j])))
	}

	sort.Ints(array1)
	sort.Ints(array2)

	return reflect.DeepEqual(array1, array2)
}
