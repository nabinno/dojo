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
