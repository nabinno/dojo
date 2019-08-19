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
