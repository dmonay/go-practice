// structs are types
package main

import "fmt"

type Counter struct {
	number int
}

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
}

// need to understand difference between make and new
// pointer and ampersand for getting location
func main() {
	// newCounter := new(Counter)

	// var anotherPointer *Counter
	// anotherPointer = newCounter

	// anotherPointer.number = 20
	// fmt.Println(newCounter.number)

	for row := range bigDigits[0] {
		fmt.Println(row)
	}
}

// interfaces can ONLY operate on non-pointers
// so must be strict type: variables are values; pointers are the address in memory --it does not return a value
// so you can't do this:
// 		func (c *Counter) Increment(someValue int) int {
// 			...
// 		}

type CounterInterfae interface {
	Increment(somevalue int) int
}

// attache it to Counter
// receivers?
// mutator methods?

func (c Counter) Increment(someValue int) int {
	c.number += someValue
	fmt.Println(c.number)
	return c.number
}
