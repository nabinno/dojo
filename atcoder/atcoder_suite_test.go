package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAtcoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Atcoder Suite")
}

var _ = Describe("Practice", func() {
	var _ = Describe("PracticeA", func() {
		It("should handle the following cases", func() {
			Expect(WelcomeAdd(1, 2, 3, "test")).To(Equal("6 test"))
			Expect(WelcomeAdd(72, 128, 256, "myonmyon")).To(Equal("456 myonmyon"))
		})
	})
})

var _ = Describe("AtCoder Beginner's Contest", func() {
	var _ = Describe("DetectHaikuLength (ABC042A)", func() {
		test := func(is [3]int, exp string) {
			Expect(DetectHaikuLength(is)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test([3]int{5, 5, 7}, "YES")
			test([3]int{7, 7, 5}, "NO")
		})
	})

	var _ = Describe("DetectHaikuLength (ABC043A)", func() {
		test := func(n int, exp int) {
			Expect(SumCandies(n)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(3, 6)
			test(10, 55)
			test(1, 1)
		})
	})

	var _ = Describe("TakeAndHotels (ABC044A)", func() {
		test := func(n, k, x, y int, exp int) {
			Expect(TakeAndHotels(n, k, x, y)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(5, 3, 10000, 9000, 48000)
			test(2, 3, 10000, 9000, 20000)
		})
	})

	var _ = Describe("Trapezoids (ABC045A)", func() {
		test := func(a, b, h, exp int) {
			Expect(Trapezoids(a, b, h)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(3, 4, 2, 7)
			test(4, 4, 4, 16)
		})
	})

	var _ = Describe("AtCoDeerAndPaintCans (ABC046A)", func() {
		test := func(a, b, c, exp int) {
			Expect(AtCoDeerAndPaintCans(a, b, c)).To(Equal(exp))
		}
		It("sould handle the following cases", func() {
			test(3, 1, 4, 3)
			test(3, 3, 33, 2)
		})
	})

	var _ = Describe("DetectEvenOrOdd (ABC086A)", func() {
		test := func(a, b int, exp string) {
			Expect(DetectEvenOrOdd(a, b)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(3, 4, "Even")
			test(1, 21, "Odd")
		})
	})
})
