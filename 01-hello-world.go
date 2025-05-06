package main

import {
	"fmt",
	"os",
}
// alternate syntax --> import "fmt" "os"

func main() {
	fmt.Println("Hello, World!")
}

/*
There are 2 ways mainly to run a Go program:
1. method1 --> go run <filename>.go
2. method2 -->
	i. go build <filename>.go (this creates the binary file)
	ii. ./<filename> (this runs the binary file)
*/