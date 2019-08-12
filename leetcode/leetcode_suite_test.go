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
