package main

import (
	"fmt"
	"strconv"
)

var i int = 13

var (
	actorName    string = "Pipi"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	// var k float32
	// k = 1
	// var i int = 42
	// j := 33
	// fmt.Println(j)
	// fmt.Println(i)
	// fmt.Println(k)
	var i int = 43
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}
