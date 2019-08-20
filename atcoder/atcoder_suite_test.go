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
	//	var _ = Describe("DetectHaikuLength (ABC042A)", func() {
	//		test := func(is []int, exp string) {
	//			Expect(DetectHaikuLength(is)).To(Equal(exp))
	//		}
	//		It("should handle the following cases", func() {
	//			test([]int{5, 5, 7}, "Yes")
	//			test([]int{7, 7, 5}, "No")
	//		})
	//	})

	var _ = Describe("DetectHaikuLength (ABC042A)", func() {
		test := func(n int, exp int) {
			Expect(SumCandies(n)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(3, 6)
			test(10, 55)
			test(1, 1)
		})
	})
})
