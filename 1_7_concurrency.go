package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	theta  float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g. 12.5 90, " +
	"or %s to quit."

// init is automatically executed before the main function and cannot be called explicitly
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+C")
	}
}

//channels are modeled on UNIX pipes and provide two-way or one-way communication (not sharing)
// of data. They behave like FIFO queues--first in, first out, preserving order.
// EX:
func channelDemo() {
	messages := make(chan string, 10) // the 10 is the buffer size. Channels in go are blocking, so if
	// the buffer fills, no messages can be sent until at least one is received

	messages <- "Leader"
	messages <- "Follower"

	// the left side of the communication operator must contain the channel and the right side
	// must be the value of the same type as the channel was declared with

	// use it in binary mode to send messages and in unary mode to receive them:
	message1 := <-messages
	message2 := <-messages

	// channels are normally created to provide communication between goroutines
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	// a go statement is given a function call that is executed asynchronously
	// in a separate goroutine. The flow of control in the current function
	// continues immiediately from the following statement (in this case the return)
	go func() {
		// infinite loop that waits until it receives a question, blocking its own goroutine
		//  (this is why it needs to be in its own async goroutine)
		for {
			polarCoord := <-questions                   // polarCoord receives the input from questions channel
			theta := polarCoord.theta * math.Pi / 180.0 // degrees to radians
			x := polarCoord.radius * math.Cos(theta)
			y := polarCoord.radius * math.Sin(theta)
			answers <- cartesian{x, y} // send converted coords to answers channel
		}
	}()
	return answers
}
