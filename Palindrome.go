package main

import (
	"fmt"
)

func main() {
	var word, revWord string
	fmt.Scan(&word)
	for _, w := range word {
		revWord = string(w) + revWord
	}
	if word == revWord {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
