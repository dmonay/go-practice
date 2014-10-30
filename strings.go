package main

import "fmt"

// a string is a slice of bytes. This is because each character is encoded with a byte, usually in UTF-8 encoding.

func main() {
	// this is represented in Unicode as U+2318
	// note the back quotes: this is a raw string, also known as a string literal
	const PlaceOfInterest = `⌘`

	// this is printed as a plain string
	fmt.Printf("plain string: ")
	fmt.Printf("%s", PlaceOfInterest)
	fmt.Printf("\n")

	// this is escaped unicode, aka ASCII-only quoted string
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", PlaceOfInterest)
	fmt.Printf("\n")

	// this is UTF-8 code, printed as individual bytes in hexadecimal 
	fmt.Printf("hex bytes: ")
	for i :=0; i < len(PlaceOfInterest); i++ {
		fmt.Printf("%x", PlaceOfInterest[i])
	}
	fmt.Printf("\n")	

	// output will be:
	// 	plain string: ⌘
	//  quoted string: "\u2318"
	//  hex bytes: e2 8c 98
}

// Go represents code points as runes. Rune is also an alias for int32. 