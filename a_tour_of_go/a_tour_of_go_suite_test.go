package main

import (
	"fmt"
	"math"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestATourOfGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ATourOfGo Suite")
}

var _ = Describe("More Types: structs, slices, and maps", func() {
	var _ = Describe("PrintSlice (moretype 11)", func() {
		testPrintSlice := func(s []int, exp string) {
			rc := PrintSlice(s)
			Expect(rc).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			testPrintSlice([]int{2, 3, 5, 7, 11, 13}, "len=6 cap=6 [2 3 5 7 11 13]\n")
		})
	})

	var _ = Describe("WordCount (moretypes 23)", func() {
		testWordCount := func(s string, key string, exp int) {
			wcMap := WordCount(s)
			Expect(wcMap[key]).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			testWordCount("hello hello hello world!!!", "hello", 3)
		})
	})

	var _ = Describe("Fibonacci (moretypes 26)", func() {
		testFibonacci := func(n int, exp []int) {
			f := Fibonacci()
			as := []int{}
			for i := 0; i < n; i++ {
				as = append(as, f())
			}
			Expect(as).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			testFibonacci(10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34})
		})
	})
})

var _ = Describe("Methods and interfaces", func() {
	var _ = Describe("Abs (methods 1)", func() {
		testAbs := func(x int, y int, exp float64) {
			v := vertex{3, 4}
			Expect(v.Abs()).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			testAbs(3, 4, 5)
		})
	})

	var _ = Describe("Scale and Abs (methods 4)", func() {
		testScale := func(x int, y int, exp float64) {
			v := vertex{3, 4}
			v.Scale(10)
			Expect(v.Abs()).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			testScale(3, 4, 50)
		})
	})

	var _ = Describe("fmt.Stringer (methods 18)", func() {
		testIPAddr := func(i ipaddr, exp string) {
			Expect(fmt.Sprintln(i)).To(Equal(exp))
		}
		It("should handle the following cases", func() {
			testIPAddr(ipaddr{127, 0, 0, 1}, "127.0.0.1\n")
		})
	})

	var _ = Describe("Sqrt (methods 20)", func() {
		testSqrtErr := func(x float64, exp interface{}) {
			if f, err := Sqrt(x); err == nil {
				Expect(f).To(Equal(exp))
			} else {
				Expect(err.Error()).To(Equal(exp))
			}
		}
		It("should handle the following cases", func() {
			testSqrtErr(float64(2), math.Sqrt(2))
			testSqrtErr(float64(-2), "cannot Sqrt negative number: -2")
		})
	})
})
