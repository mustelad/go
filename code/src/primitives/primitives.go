package main

import (
	"fmt"
)

func main() {
	var n bool = true
	i := 1 == 1
	j := 1 == 2
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", n, n)
	var o uint16 = 42
	fmt.Printf("%v, %T\n", o, o)
}
