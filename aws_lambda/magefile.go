//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	lbd "github.com/nabinno/dojo/aws_lambda/magefiles/aws_lambda"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Lbd ________________________________________
type Lbd mg.Namespace

// Dispatch is command that dispatch on local
func (Lbd) Dispatch() { lbd.DispatchOnLocal() }

// Test is command that test
func (Lbd) Test() { lbd.TestOnLocal() }

// Build is command that build
func (Lbd) Build() { lbd.Build() }

// Deploy is command that deploy
func (Lbd) Deploy() { lbd.Deploy() }

// Clean is command that clean
func (Lbd) Clean() { lbd.Clean() }
