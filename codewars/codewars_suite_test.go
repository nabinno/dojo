package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCodewars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Codewars Suite")
}

var _ = Describe("7kyu", func() {
	var _ = Describe("NbDig (Count the Digit)", func() {
		testNgDig := func(n int, d int, exp int) {
			var ans = NbDig(n, d)
			Expect(ans).To(Equal(exp))
		}
		It("should handle basic cases", func() {
			testNgDig(550, 5, 213)
			testNgDig(5750, 0, 4700)
		})
	})

	var _ = Describe("Basic tests", func() {
		It("should return the correct values", func() {
			Expect(GenerateBandName(string("knife"))).Should(BeEquivalentTo("The Knife"))
			Expect(GenerateBandName(string("tart"))).Should(BeEquivalentTo("Tartart"))
			Expect(GenerateBandName(string("sandles"))).Should(BeEquivalentTo("Sandlesandles"))
			Expect(GenerateBandName(string("bed"))).Should(BeEquivalentTo("The Bed"))
		})
	})
})

var _ = Describe("6kyu", func() {
	var _ = Describe("MaxBall", func() {
		test := func(v0 int, exp int) {
			var ans = MaxBall(v0)
			Expect(ans).To(Equal(exp))
		}
		It("should handle basic cases", func() {
			test(37, 10)
			test(45, 13)
			test(99, 28)
			test(85, 24)
			test(136, 39)
		})
	})

	var _ = Describe("ValidBraces", func() {
		test := func(str string, exp bool) {
			Expect(ValidBraces(str)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test("(){}[]", true)
			test("([{}])", true)
			test("(}", false)
			test("[(])", false)
			test("[({})](]", false)
		})
	})
})
