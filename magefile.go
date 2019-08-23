//+build mage

package main

import (
	"github.com/magefile/mage/mg"
	lambda "github.com/nabinno/dojo/magefiles/aws_lambda"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Lambda ________________________________________
type Lambda mg.Namespace

// Pack ....
func (Lambda) Pack() { lambda.Pack() }

// Build ....
func (Lambda) Build() { lambda.Build() }

// InstallDeps ....
func (Lambda) InstallDeps() { lambda.InstallDeps() }

// Clean ....
func (Lambda) Clean() { lambda.Clean() }
