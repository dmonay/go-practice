package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	sliceOfWords := strings.Fields(s)
	myMap := make(map[string]int)
	mapOfTicks := make(map[int]int)
	for i := range sliceOfWords {
		curWord := sliceOfWords[i]
		// the inner loop does the comparison
		for b := range sliceOfWords {
			if sliceOfWords[b] == curWord {
				mapOfTicks[i]++
			}
		}
	}

	for z, word := range sliceOfWords {
		myMap[word] = mapOfTicks[z]
	}
	return myMap
}

func main() {
	fmt.Println(WordCount("hey hey hey you are you"))
}
