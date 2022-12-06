package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("Hello World \n OS Arguments: %v \n Argument 1: %v \n Argument 1 and beyond: %v", args, args[1], args[1:])
}
