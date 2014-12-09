package main

import "fmt"

func f1(c chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("I'm doing stuff")
	c <- "done!"
}

// run 10 gorountines
func main() {
	done := make(chan string)
	go f1(done)

	msg := <-done
	fmt.Println(msg)
}
