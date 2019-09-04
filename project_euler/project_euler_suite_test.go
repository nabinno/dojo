package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProjectEuler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ProjectEuler Suite")
}

var _ = Describe("GetLargestPrimeFactorOf (problem 3)", func() {
	test := func(n int64, exp int64) {
		Expect(GetLargestPrimeFactorOf(n)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(600851475143, 6857)
	})
})

var _ = Describe("GetLargestPalindromeProduct (problem 4)", func() {
	test := func(n int, exp string) {
		Expect(GetLargestPalindromeProduct(n)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(2, "9009 = 91 x 99")
		test(3, "906609 = 913 x 993")
	})
})

var _ = Describe("GetSmalletsMultiple (problem 5)", func() {
	test := func(n int, exp int) {
		Expect(GetSmalletsMultiple(n)).To(Equal(exp))
	}
	It("should handle the follwoing cases", func() {
		test(10, 2520)
		test(20, 232792560)
	})
})

var _ = Describe("SumSquareDifference (problem 6)", func() {
	test := func(n int, exp int) {
		Expect(SumSquareDifference(n)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(10, 2640)
		test(100, 25164150)
	})
})

var _ = Describe("Get10001stPrime (problem 7)", func() {
	test := func(nth, exp int) {
		Expect(GetNthPrime(nth)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(10001, 104743)
	})
})
