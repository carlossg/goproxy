package main

import (
	"fmt"

	"github.com/Microsoft/hcsshim"
)

func main() {
	fmt.Println("Hello, World!")
	hcsshim.CreateComputeSystem("test", hcsshim.ComputeSystemCreateOptions{})
}
