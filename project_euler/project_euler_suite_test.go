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

var _ = Describe("GetLargestPalindromeProduct (problem 4)", func() {
	test := func(n int, exp string) {
		Expect(GetLargestPalindromeProduct(n)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(2, "9009 = 91 x 99")
		test(3, "906609 = 913 x 993")
	})
})

var _ = Describe("GetSmalletsMultiple (problem 5)", func() {
	test := func(n int, exp int) {
		Expect(GetSmalletsMultiple(n)).To(Equal(exp))
	}
	It("should handle the follwoing cases", func() {
		test(10, 2520)
		test(20, 232792560)
	})
})

var _ = Describe("SumSquareDifference (problem 6)", func() {
	test := func(n int, exp int) {
		Expect(SumSquareDifference(n)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(10, 2640)
		test(100, 25164150)
	})
})

var _ = Describe("Get10001stPrime (problem 7)", func() {
	test := func(nth, exp int) {
		Expect(GetNthPrime(nth)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(10001, 104743)
	})
})

var _ = Describe("GetLargestProductInASeries (problem 8)", func() {
	test := func(digit string, exp int) {
		Expect(GetLargestProductInASeries(digit)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test("7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450", 1)
	})
})

var _ = Describe("GetSpecialPythagoreanTriplet (problem 9)", func() {
	test := func(sumInt, exp int) {
		Expect(GetSpecialPythagoreanTriplet(sumInt)).To(Equal(exp))
	}
	It("should handle the following cases", func() {
		test(1000, 1)
	})
})
