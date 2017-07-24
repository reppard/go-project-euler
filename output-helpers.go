package main

import (
	"fmt"
)

const borderString = "//"

func printBorderedMsg(message []string) {
	printBar()
	for i := 0; i < len(message); i++ {
		fmt.Println(borderString, message[i])
	}
	printBar()
	fmt.Println("\n")
}

func printBar() {
	s := ""

	for i := 1; i <= 40; i++ {
		s += borderString
	}
	fmt.Println(s)
}
