//+build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/mattn/go-shellwords"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

// Atcoder ____________________
type Atcoder mg.Namespace

// Test ....
func (Atcoder) Test() {
	fmt.Println("Testing...")
	cmd("ginkgo atcoder")
}

// Atourofgo ____________________
type Atourofgo mg.Namespace

// Test ....
func (Atourofgo) Test() {
	fmt.Println("Testing...")
	cmd("ginkgo a_tour_of_go")
}

// Codewars ____________________
type Codewars mg.Namespace

// Test ....
func (Codewars) Test() {
	fmt.Println("Testing...")
	cmd("ginkgo codewars")
}

// Leetcode ____________________
type Leetcode mg.Namespace

// Test ....
func (Leetcode) Test() {
	fmt.Println("Testing...")
	cmd("ginkgo leetcode")
}

// Projecteuler ____________________
type Projecteuler mg.Namespace

// Test ....
func (Projecteuler) Test() {
	fmt.Println("Testing...")
	cmd("ginkgo project_euler")
}

// Awslambda ____________________
type Awslambda mg.Namespace

const coreBinaryName = "aws-lambda-go-api-proxy-core"
const ginBinaryName = "aws-lambda-go-api-proxy-gin"
const sampleBinaryName = "main"

// Pack ....
func (Awslambda) Pack() {
	fmt.Println("Packing...")
	cmd(fmt.Sprintf("cd sample && zip main.zip %v", sampleBinaryName))
}

// Test ....
func (Awslambda) Test() {
	fmt.Println("Testing...")
	cmd("ginkgo aws_lambda")
}

// Build step that requires additional params, or platform specific steps for example
func (Awslambda) Build() {
	mg.Deps(Awslambda.InstallDeps)
	fmt.Println("Building...")
	cmd(fmt.Sprintf("go build ./... && cd aws_lambda && go build -o %v", sampleBinaryName))
}

// InstallDeps manage your deps, or running package managers.
func (Awslambda) InstallDeps() {
	fmt.Println("Installing Deps...")
	cmd("go get github.com/stretchr/piglatin")
}

// Clean up after yourself
func (Awslambda) Clean() {
	fmt.Println("Cleaning...")
	_ = os.RemoveAll(fmt.Sprintf("sample/%v*", sampleBinaryName))
}

func cmd(cmdstr string) {
	if out, err := doCmd(cmdstr); err == nil {
		fmt.Println(string(out))
	} else {
		fmt.Println(err)
	}
}

func doCmd(cmdstr string) (out []byte, err error) {
	args, err := shellwords.Parse(cmdstr)
	if err != nil {
		return
	}

	switch len(args) {
	case 0:
		return
	case 1:
		out, err = exec.Command(args[0]).Output()
	default:
		out, err = exec.Command(args[0], args[1:]...).Output()
	}
	if err != nil {
		return
	}
	return
}
