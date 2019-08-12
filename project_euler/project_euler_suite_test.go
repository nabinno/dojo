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
