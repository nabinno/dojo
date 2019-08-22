package main

import (
	"fmt"
	"math"

	"github.com/thoas/go-funk"
	"github.com/wesovilabs/koazee"
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

	key := koazee.
		StreamOf(funk.Keys(rc)).
		Sort(func(x, y int) int {
			if x <= y {
				return -1
			}
			return 1
		}).
		Last().
		Int()

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
