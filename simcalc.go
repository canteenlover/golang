package main

import (
	"fmt"
)

func main() {
	var a, b int
	var symb string
	fmt.Print("Input 1st num: ")
	fmt.Scan(&a)
	fmt.Print("Input 2nd num: ")
	fmt.Scan(&b)
	fmt.Print("Input one of the characters from the list (+ - * /): ")
	fmt.Scan(&symb)

	switch {
	case symb == "+":
		fmt.Printf("%d %s %d = %d", a, symb, b, a+b)
	case symb == "-":
		fmt.Printf("%d %s %d = %d", a, symb, b, a-b)
	case symb == "*":
		fmt.Printf("%d %s %d = %d", a, symb, b, a*b)
	case symb == "/":
		fmt.Printf("%d %s %d = %d", a, symb, b, a/b)
	}
}
