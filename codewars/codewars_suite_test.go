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

var _ = Describe("NbDig (7kyu Count the Digit)", func() {
	testNgDig := func(n int, d int, exp int) {
		var ans = NbDig(n, d)
		Expect(ans).To(Equal(exp))
	}

	It("should handle basic cases", func() {
		testNgDig(550, 5, 213)
		testNgDig(5750, 0, 4700)
	})
})
