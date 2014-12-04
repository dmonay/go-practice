package main

import (
	"fmt"
)

func main() {
	// this can be called with any number of args
	fmt.Println(add(1, 2, 3))

	// when it's called, value of local var i persists:
	// it persists b/c the internal function is being assigned to nextEven, which is one slot in memory
	// thus each time it's called, the scope remains, it isn't reset.
	// this takes up more memory but is more performant
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2
	fmt.Println(nextEven()) // 4

	// if it's executed as an immediate function, var i does NOT persist
	// this is b/c it's an anonymous function so the reference to the variable is lost
	// after it's executed, the garabage collector clears the spot in memory that it took up
	// thus it's less memory intensive, but also less performant than the above implementation, b/c
	// it has to be executed anew every time.
	fmt.Println(makeEvenGenerator()()) // 0
	fmt.Println(makeEvenGenerator()()) // 0
	fmt.Println(makeEvenGenerator()()) // 0

	// fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// ****************** FUNCTIONS **********************

// a signature is the parameter and the return type, collectively

// a function can return multiple values:

func f() (int, int) {
	return 5, 6
}

// variadic functions
// the elipses before the type name of the last param indicates that it takes 0 or more of those params
func add(args ...int) int {
	total := 0
	for _, value := range args {
		total += value
	}
	return total
}

// *************** CLOSURES **************

func makeEvenGenerator() func() uint {
	i := uint(0) // this is not garbage collected, because it remains in scope, and thus is persists between calls!!!
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// fibonacci is a function that returns a function that returns an int
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

// niladic functions are functions without input values. The init() function is such.

// ********************** DEFER, PANIC, RECOVER *****************
// the defer statement schedules a function call to be used after a function is complete
// output is 1, then 2
func sequential() {
	defer fmt.Println(2)
	fmt.Println(1)
}

// recover is a statement that stops a panic and returns the value that was passed to the call to panic
// it must be used inside a defer statement because panic() stops the execution of a function:
func main2() {
	// induce a runtime error:
	panic("oh no!")
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
}

// a composite literal is an expression that creates a new value each time it is evaluated. E.g. creating a new struct:

type testRectangle struct {
	length int
	width  int
}

var rect1 = testRectangle{10, 5}
