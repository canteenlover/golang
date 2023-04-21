package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
	"unicode"
)

/*This code checks if your string is "right"
* Right means that your string starts with Capital letter and ends with dot "."
 */

func main() {
	var pref, suff bool
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	textRunes := []rune(text)
	if unicode.IsUpper(textRunes[0]) {
		pref = true
	}

	text = strings.TrimSuffix(text, "\r\n")

	if strings.HasSuffix(text, ".") {
		suff = true
	}

	if pref == true && suff == true {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
