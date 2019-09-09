package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLeetcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Leetcode Suite")
}

var _ = Describe("700 ~ 799", func() {
	var _ = Describe("To Lower Case (709)", func() {
		test := func(s string, exp string) {
			Expect(ToLowerCase(s)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test("Hello", "hello")
			test("here", "here")
			test("LOVELY", "lovely")
		})
	})

	var _ = Describe("NumJewelsInStones (771)", func() {
		test := func(j string, s string, exp int) {
			Expect(NumJewelsInStones(j, s)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test("aA", "aAAbbbb", 3)
		})
	})
})

var _ = Describe("800 ~ 899", func() {
	var _ = Describe("Unique Morse Code Words (804)", func() {
		test := func(words []string, exp int) {
			Expect(UniqueMorseRepresentations(words)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test([]string{"gin", "zen", "gig", "msg"}, 2)
		})
	})

	var _ = Describe("Flipping an Image (832)", func() {
		test := func(A [][]int, exp [][]int) {
			Expect(FlipAndInvertImage(A)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			test([][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}}, [][]int{[]int{1, 0, 0}, []int{0, 1, 0}, []int{1, 1, 1}})
		})
	})
})

var _ = Describe("SortArrayByParity (905)", func() {
	test := func(A, exp []int) {
		Expect(SortArrayByParity(A)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test([]int{3, 1, 2, 4}, []int{2, 4, 3, 1})
	})
})

var _ = Describe("RangeSumBST (938)", func() {
	test := func(root *treeNode, L int, R int, exp int) {
		Expect(RangeSumBST(root, L, R)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(convertIntsToTreeNode([]int{10, 5, 15, 3, 7, null, 18}), 7, 15, 32)
		test(convertIntsToTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}), 6, 10, 23)
	})
})

var _ = Describe("Remove Outermost Parentheses (1021)", func() {
	test := func(S string, exp string) {
		Expect(RemoveOuterParentheses(S)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test("(()())(())", "()()()")
		test("(()())(())(()(()))", "()()()()(())")
		test("()()", "")
	})
})

var _ = Describe("DefangIPaddr (1108)", func() {
	test := func(a string, exp string) {
		Expect(DefangIPaddr(a)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test("0.0.0.0", "0[.]0[.]0[.]0")
	})
})
