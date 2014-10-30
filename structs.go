package main

import (
	"fmt"
	"math"
)

// ************* STRUCTS *****************************
type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := r.x1 * r.y1
	w := r.x2 * r.y2
	return l * w
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c)) // call as a function --> must use &
	fmt.Println(c.area())       // call as a method --> dot notation possible b/c pointer type used when method was defined

	// call Android
	a := new(Android)
	a.Name = "John" // or a.Person.Name
	a.Person.Talk() // or call directly: a.Talk()
}

// we can turn the function circleArea into a method of the struct Circle
// area() is the method and can be called with the dot notation
// this is the syntax for creating a method:
// c is the receiver and *Circle is the type of the receiver, in this case it's a pointer to Circle, a custom type
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// ************* EMBEDDED TYPES *****************************

// structs are normally "has-a" relationships, but it's possinle to create "is-a" relationship using embedded types:

// person struct
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

// android struct where Android HAS a person
// type Android struct {
// 	Person Person
// 	Model string
// }

// better struct wehere an android IS a person
type Android struct {
	Person
	Model string
}

// ************* INTERFACES *****************************

// rectangle and circle both have the method area(). Thats' because they are both shapes. This method can be defined in an interface, which is a set of methods:

type Shape interface {
	area() float64
}

// this is useful because interfaces can be used as arguments to functions
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// the ellipsis before the type name of the last param indicates that it takes 0 or more of those params
// the underscore denotes a variable that is declared but not used because the Go compiler won't allow you to create variables you don't use
// this is then called like this:
// fmt.Println(totalArea(&c, &r))
