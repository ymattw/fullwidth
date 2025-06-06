// Find all fullwidth unicode codepoints.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

func main() {
	const maxRune = utf8.MaxRune
	count := 0

	for r := rune(0); r <= maxRune; r++ {
		if r >= 0xD800 && r <= 0xDFFF {
			continue // skip UTF-16 surrogate halves
		}
		if !unicode.IsGraphic(r) {
			continue // Skip non visiable ones
		}

		ch := string(r)
		if uniseg.StringWidth(ch) != 2 {
			continue
		}

		fmt.Printf(" %s", ch)
		count++
		if count%50 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
