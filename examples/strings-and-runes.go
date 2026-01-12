package examples

import (
	"fmt"
	"unicode/utf8"
)

func exanimeRune(r rune) {
	switch r {
	case 't':
		fmt.Println("found tee")
	case 'ส':
		fmt.Println("found so sua")
	}
}

func StringsAndRunes() {
	const s = "สวัสดี"
	fmt.Println("Len:", len(s))

	for i := range len(s) {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for i, r := range s {
		fmt.Printf("%#U starts at %d\n", r, i)
	}

	fmt.Println("Using utf8 package to decode:")
	for i, w := 0, 0; i < len(s); i += w {
		r, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", r, i)
		w = width
		exanimeRune(r)
	}
}
