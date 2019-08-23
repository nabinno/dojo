package awslambda

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/nabinno/dojo/magefiles/utilities"
)

var cmd = utilities.Cmd

const coreBinaryName = "aws-lambda-go-api-proxy-core"
const ginBinaryName = "aws-lambda-go-api-proxy-gin"
const sampleBinaryName = "main"

// Pack ....
func Pack() {
	fmt.Println("Packing...")
	if err := os.Chdir("aws_lambda"); err == nil {
		cmd(fmt.Sprintf("zip main.zip %v", sampleBinaryName))
	}
}

// Build step that requires additional params, or platform specific steps for example
func Build() {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd("go build ./...")
	if err := os.Chdir("aws_lambda"); err == nil {
		cmd(fmt.Sprintf("go build -o %v", sampleBinaryName))
	}
}

// InstallDeps manage your deps, or running package managers
func InstallDeps() {
	fmt.Println("Installing Deps...")
	cmd("go get github.com/stretchr/piglatin")
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	if err := os.RemoveAll(fmt.Sprintf("aws_lambda/%v*", sampleBinaryName)); err == nil {
		fmt.Println("Success cleaning!")
	}
}
