package main

import (
	"fmt"
	"strings"
)

//remove all characters that occur more than once

func main() {
	var word, res string
	fmt.Scan(&word)
	for _, i := range word {
		if strings.Count(word, string(i)) == 1 {
			res += string(i)
		}
	}
	fmt.Println(res)
}
