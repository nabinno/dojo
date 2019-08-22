package main

import (
	"fmt"
	"math"
	"sort"
)

// GetLargestPalindromeProduct ...
func GetLargestPalindromeProduct(digitNumber int) string {
	var i, j int
	rc := make(map[int][2]int)
	min := int(math.Pow10(digitNumber-1)) - 1
	max := int(math.Pow10(digitNumber)) - 1

	for i = max; i > min; i-- {
		for j = max; j > min; j-- {
			if isPalindrome(i * j) {
				rc[i*j] = [2]int{i, j}
			}
		}
	}

	key := getMaxInit(getKeys(rc))

	return fmt.Sprintf("%v = %v x %v", key, rc[key][0], rc[key][1])
}

func isPalindrome(n int) bool {
	if n < 0 || (n%10 == 0 && n != 0) {
		return false
	}

	revertedNum := 0
	for n > revertedNum {
		revertedNum = revertedNum*10 + (n % 10)
		n = n / 10
	}
	return n == revertedNum || n == revertedNum/10
}

// @see https://stackoverflow.com/questions/21362950/getting-a-slice-of-keys-from-a-map
func getKeys(m map[int][2]int) (keys []int) {
	keys = make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return
}

func getMaxInit(is []int) int {
	sort.Ints(is)
	return is[len(is)-1]
}
