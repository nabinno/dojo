package awslambda

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

// Build step that requires additional params, or platform specific steps for example
func Build() {
	mg.Deps(doBuild, doPack)
}

func doBuild() {
	fmt.Println("Building...")
	if err := os.Chdir("aws_lambda"); err == nil {
		cmd("go build ./...")
		cmd(fmt.Sprintf("go build -o %v", SampleBinaryName))
	}
}

func doPack() {
	fmt.Println("Packing...")
	if err := os.Chdir("aws_lambda"); err == nil {
		cmd(fmt.Sprintf("zip main.zip %v", SampleBinaryName))
	}
}
