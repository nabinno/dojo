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

var _ = Describe("RangeSumBST (938)", func() {
	test := func(root *treeNode, L int, R int, exp int) {
		Expect(RangeSumBST(root, L, R)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(convertIntsToTreeNode([]int{10, 5, 15, 3, 7, null, 18}), 7, 15, 32)
		test(convertIntsToTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, null, 6}), 6, 10, 23)
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

var _ = Describe("NumJewelsInStones (771)", func() {
	test := func(j string, s string, exp int) {
		Expect(NumJewelsInStones(j, s)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test("aA", "aAAbbbb", 3)
	})
})
