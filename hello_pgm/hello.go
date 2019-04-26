package main

import (
	"fmt"
	"math"
	"math/rand"
)

var aa, bb, cc bool = true, false, true

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello there Programmer person")
	fmt.Println("A random number is ", rand.Intn(100))
	fmt.Println("Pi ", math.Pi)
	fmt.Println("An add function ", add(63, 27))
	dd := 42
	fmt.Println(aa, bb, cc, dd)
}

