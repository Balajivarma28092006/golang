package main

import (
	"fmt"
	"os"
)

func main() {
	data := os.Args[1]
	fmt.Printf("Arg passed was: %v", data)
}
