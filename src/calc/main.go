package main

import (
	"fmt"
	i "github.com/skilstak/go-input"
	"strconv"
	s "strings"
)

func add(x, y float64) float64 {
	return x + y
}

func sub(x, y float64) float64 {
	return x - y
}

func mult(x, y float64) float64 {
	return x * y
}

func div(x, y float64) float64 {
	return x / y
}

func main() {
	fmt.Println("[a] Addition")
	fmt.Println("[s] Subtraction")
	fmt.Println("[m] Multiplication")
	fmt.Println("[d] Division")
	option, _ := i.Prompt("> ")
	num1, _ := i.Prompt("First number: ")
	num2, _ := i.Prompt("Second number: ")

	flt1, _ := strconv.ParseFloat(num1, 64)
	flt2, _ := strconv.ParseFloat(num2, 64)

	if s.Contains(option, "a") {
		fmt.Println(add(flt1, flt2))
	} else if s.Contains(option, "s") {
		fmt.Println(sub(flt1, flt2))
	} else if s.Contains(option, "m") {
		fmt.Println(mult(flt1, flt2))
	} else if s.Contains(option, "d") {
		fmt.Println(div(flt1, flt2))
	}
}
