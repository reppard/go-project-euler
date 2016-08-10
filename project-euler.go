package main

import (
	"flag"
	"fmt"
)

func problem1() {
	fmt.Println("///////////////////////////////////////////////////////////////////////////////")
	fmt.Println("// Problem 1:")
	fmt.Println("// If we list all the natural numbers below 10 that are multiples of 3 or 5,")
	fmt.Println("// we get 3, 5, 6 and 9. The sum of these multiples is 23.")
	fmt.Println("//")
	fmt.Println("///////////////////////////////////////////////////////////////////////////////")
	fmt.Println("\nFind the sum of all the multiples of 3 or 5 below 1000.")

	sum := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("Answer: %d\n", sum)
}

func main() {
	fmt.Println("///////////////////////////////////////////////////////////////////////////////")
	fmt.Println("//")
	fmt.Println("// Project Euler in Go")
	fmt.Println("//")
	fmt.Println("///////////////////////////////////////////////////////////////////////////////")

	problemPtr := flag.String("problem", "error", "Specify a problem to execute (default: problem1)")

	flag.Parse()

	problem := *problemPtr

	switch problem {
	case "problem1":
		problem1()
	default:
		problem1()
	}
}
