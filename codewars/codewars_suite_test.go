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

	var _ = Describe("CartesianNeighbor", func() {
		test := func(x, y int, exp [][]int) {
			Expect(CartesianNeighbor(x, y)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(2, 2, [][]int{[]int{1, 1}, []int{1, 2}, []int{1, 3}, []int{2, 1}, []int{2, 3}, []int{3, 1}, []int{3, 2}, []int{3, 3}})
		})
	})

	var _ = Describe("Comp", func() {
		test := func(s1, s2 []int, exp bool) {
			Expect(Comp(s1, s2)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(
				[]int{121, 144, 19, 161, 19, 144, 19, 11},
				[]int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
				true,
			)
		})
		It("should handle the following cases", func() {
			test(
				[]int{121, 144, 19, 161, 19, 144, 19, 11},
				[]int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
				false,
			)
		})
		It("should handle the following cases", func() {
			test(
				nil,
				[]int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
				false,
			)
			test(nil, nil, true)
		})
		It("should handle the following cases", func() {
			test(
				[]int{121, 144, 19, 161, 19, 144, 19, 11},
				[]int{132, 14641, 20736, 361, 25921, 361, 20736, 361},
				false,
			)
		})
	})

	var _ = Describe("InArray", func() {
		test := func(array1, array2 []string, exp []string) {
			Expect(InArray(array1, array2)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(
				[]string{"arp", "live", "strong"},
				[]string{"lively", "alive", "harp", "sharp", "armstrong"},
				[]string{"arp", "live", "strong"},
			)
			test(
				[]string{"tarp", "mice", "bull"},
				[]string{"lively", "alive", "harp", "sharp", "armstrong"},
				[]string{},
			)
		})
	})

	var _ = Describe("Wave", func() {
		test := func(words string, exp []string) {
			Expect(Wave(words)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(" x yz", []string{" X yz", " x Yz", " x yZ"})
		})
	})
})

var _ = Describe("5kyu", func() {
	var _ = Describe("ProductFib", func() {
		test := func(prod uint64, exp [3]uint64) {
			Expect(ProductFib(prod)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test(4895, [3]uint64{55, 89, 1})
			test(5895, [3]uint64{89, 144, 0})
			test(74049690, [3]uint64{6765, 10946, 1})
			test(84049690, [3]uint64{10946, 17711, 0})
		})
	})
})
