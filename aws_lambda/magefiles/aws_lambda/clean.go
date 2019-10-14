package awslambda

import (
	"fmt"
	"os"
)

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	if err := os.RemoveAll(fmt.Sprintf("aws_lambda/%v*", SampleBinaryName)); err == nil {
		fmt.Println("Success cleaning!")
	}
}
