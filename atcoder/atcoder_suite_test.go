package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAtcoder(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Atcoder Suite")
}

// var _ = Describe("AtCoder ABC", func() {
//	var _ = Describe("DetectHaikuLength (ABC042A)", func() {
//		test := func(is []int, exp string) {
//			Expect(DetectHaikuLength(is)).To(Equal(exp))
//		}
//		It("should handle the following cases", func() {
//			test([]int{5, 5, 7}, "Yes")
//			test([]int{7, 7, 5}, "No")
//		})
//	})
// })
