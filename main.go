package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

func main() {
	tests := []string{
		"\uFF76",
		"\uFF76\uFF9E",
		"\uFF9E",
		"\U0001F9D1",
		"\U0001F9D1\u200D\U0001F33E",
		"\U0001F33E",
	}

	for i, s := range tests {
		fmt.Printf("Case %d\n", i+1)
		fmt.Printf("String: `%v` (%#v)\n", s, s)
		fmt.Printf("StringWidth: %d\n", runewidth.StringWidth(s))
		rw := 0
		for _, c := range s {
			rw += runewidth.RuneWidth(c)
		}
		fmt.Printf("Sum of RuneWidth: %d\n", rw)
		fmt.Println()
	}
}
