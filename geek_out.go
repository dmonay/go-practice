package main

import "fmt"

type Counter struct {
	number int
}

type Counter2 struct {
	blah int
}

type CounterInterface interface {
	Increment(someValue int) int
}

func (c Counter) Increment(someValue int) int {
	c.number += someValue
	fmt.Println(c.number)
	return c.number
}

func (c Counter2) Increment(someValue int) int {
	c.blah += someValue
	fmt.Println(c.number)
	return c.number
}

func IncreaseBy10(c CounterInterface) {
	c.Increment(10)
}

func SetNumberTo20(c *Counter) {
	c.number = 20
}

func main() {

	counter := Counter{}
	IncreaseBy10(counter)       // 10 calling Increment method
	fmt.Println(counter.number) // 0

	SetNumberTo20(&counter)
	fmt.Println(counter.number) // 20

	// different instances
	// newCounter := Counter{}
	// blahCounter := newCounter

	// someInterface := interface{}{}

	// blahCounter.number = 10
	// fmt.Println(newCounter.number)

	// pointer instances
	// newCounter := new(Counter)

	// var anotherPointer *Counter
	// anotherPointer = newCounter

	// anotherPointer.number = 10
	// fmt.Println(newCounter.number) // 10

	// newCounter.number = 20
	// fmt.Println(anotherPointer.number) // 20

	// newCounter := Counter{}
	// fmt.Println(newCounter.number)
}
