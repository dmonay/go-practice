package main

import "fmt"

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // still 5

	zeroPointer(&x)
	fmt.Println(x) // now set to 0
}

// pointers reference the location of a variable in memory, not the variable's value
// this function will NOT modify the value of its argument b/c of scope
func zero(x int) {
	x = 0
}

// this function uses a pointer and WILL modify value
func zeroPointer(xPtr *int) {
	*xPtr = 0 // this says store the int 0 in the memory location *xPtr refers to
}

// the * operator is used to represent a pointer, followed by type of stored value
// it's also used to dereference pointer variables - this gives us access to the VALUE
// the pointer points to

// the & operator refers to the address of a variable
