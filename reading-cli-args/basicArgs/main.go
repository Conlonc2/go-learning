package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	fmt.Println(args)

	for i, val := range args {
		fmt.Printf("Arg %d: %v\n", i, val)
	}
	fmt.Printf("Number of Args: %v", len(args))
}
