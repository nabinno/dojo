//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	lbd "github.com/nabinno/dojo/magefiles/aws_lambda"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Lbd ________________________________________
type Lbd mg.Namespace

// Dispatch ....
func (Lbd) Dispatch() { lbd.DispatchOnLocal() }

// Test ....
func (Lbd) Test() { lbd.TestOnLocal() }

// Build ....
func (Lbd) Build() { lbd.Build() }

// Deploy ....
func (Lbd) Deploy() { lbd.Deploy() }

// Clean ....
func (Lbd) Clean() { lbd.Clean() }
