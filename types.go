package main

import (
	"fmt"
)

// ****************** ARRAYS **********************

// create an array of 5 integers
var x [5]int

// populate it quickly:
// x := [5]float64{ 98, 93, 77, 82, 83 }

func main() {
	x := [5]float64{98, 93, 77, 82, 83}
	x[4] = 100 // set the fifth element to 100
}

// this is why we use the underscore:
// this function only grabs the index of the array and adds up the indices
func getIndex() {
	x := []int{1, 2, 97} // I used a slice here cause I didn't know ahead of time how many elements my array would have
	for index := range x {
		total += index
	}
	fmt.Println(total) // 0 + 1 + 2 = 3
}

// if you want to actually grab the value at that index, use the undescore:

func getValue() {
	x := []int{1, 2, 97}
	for _, value := range x {
		total += value
	}
	fmt.Println(total) // 1 + 2 + 97 = 100
}

// ****************** SLICES **********************

// have two built in methods:

func appendSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5) // [1,2,3,4,5]
}

// other method is copy

// use make() to initialize a slice:
var pow = make([]int, 10) // this allocates a zeroed array of length  10. To specify capacity, pass in a third argument
// make is one of two allocation primitives

// ************ MAPS/HASHES/DICTIONARIES *******************

// maps are unordered collections of key-value pairs
// map with key type string and value type int
// this declares the map but doesn't initialize it
var myMap map[string]int


// to initialize it:
func makeMap() {
	myRealMap := make(map[string]int)
	x["key"] = 10         // set the key "key" to value 10
	fmt.Println(x["key"]) // use the same syntax as defining the key to access the key

	//shorter way:
	shortMap := {
		"H": "Hydrogen",
		"He": "Helium"
	}
}


func checkMap() {
	m := make(map[string]string)
	m["name"] = "John"

	// delete a key:
	delete(m, "name")

	// check that a key is present:
	v, ok := m["name"] // ok is false
}



