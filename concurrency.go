package main

import (
	"fmt"
	"math/rand"
	"time"
)

// c can only be sent to
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// channel is not restricted (can send and receive). This is known as
// bi-directional channel
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		// time.Sleep(time.Second * 1)
	}
}

func finish(c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
func main() {
	c := make(chan string)
	done := make(chan bool)

	go pinger(c)
	go ponger(c)
	go printer(c)
	go finish(done)
	<-done
}

// the select statement is like a switch statement but for channels.

// select {
// case msg1 := <- c1:
//     fmt.Println("Message 1", msg1)
// case msg2 := <- c2:
//     fmt.Println("Message 2", msg2)
// case <- time.After(time.Second):
//     fmt.Println("timeout")
// }

// if you pass in a third argument to the make statement, you create a buffered channel.
// Normally channels are synchronous. Buffered channels are asynchronous. Sending or receiving
// a message will not wait unless the channel is full.

// concurrency is NOT parallelism. Concurrency is the structure that may make parallelism possible - it is the way
// you set up the processes and synchronize them. Concurrency is the composition of independently executing processes.
// Parallelism is the simultaneous execution of processes.
