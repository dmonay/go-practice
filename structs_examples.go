package main

import "fmt"

// we use pointers when we want to modify the original value. for example, the length method just gets the length of a string so it's fine to operate on a copy of a struct
// however the Pop() method modifies an array, so we'd use a pointer to make sure our changes are recorded on the original struct and not on a copy:

// type myString interface{}

// func (stringRec myString) myLen() {
// 	return len(stringRec)
// }

// func (otherStringRec *myString) addSpace() {
// 	*otherStringRec = *otherStringRec + " spaced"
// }

// func main() {
// 	m := myString
// }

type mySlice []interface{}

func (ms mySlice) myLen() int {
	return len(ms)
}

func (ms *mySlice) myPush(x interface{}) {
	*ms = append(*ms, x)
}

func main() {
	var m mySlice
	m.myPush("job")
	fmt.Println(m)
}
