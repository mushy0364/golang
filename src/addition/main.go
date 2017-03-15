package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	x := 20
	y := 20
	fmt.Println(add(x, y))
}
