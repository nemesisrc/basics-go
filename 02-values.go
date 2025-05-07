// value types: int, float, string, bool
// reference types: array, slice, map, struct, channel, function, interface

package main

import "fmt"

func main() {
	fmt.Println("hello" + " " + "raj")

	fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}

