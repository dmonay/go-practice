package main

import (
	"fmt"
	"os/exec"
	"time"
)

// func pomodoroTurn(c chan string) {
// 	turn := 1
// 	time.Sleep(time.Second * 10)
// 	c <- "One pomodoro turn done!"
// 	turn++

// }

// func shortBreak(c chan string) {
// 	time.Sleep(time.Second * 5)
// 	c <- "Short break has ended!"
// }

// func main() {
// 	chan1 := make(chan string)
// 	fmt.Println("\x1b[32;1mYou've launched the Pomodoro!\x1b[0m")
// 	go pomodoroTurn(chan1)
// 	go shortBreak(chan1)
// 	msg := <-chan1
// 	fmt.Println(msg)

// }

var (
	currentTurn = 1
	totalTurns  = 3
)

func main() {
	turn := make(chan bool)
	smallBreak := make(chan bool)
	longBreak := make(chan bool)
	done := make(chan bool)

	go pomodoroTurn(turn)
	go pomodoroService(turn, smallBreak, longBreak, done)

	<-done
}

func pomodoroTurn(chanPomodoro chan bool) {
	tellBeginTurn()
	// time.Sleep(time.Minute * 25)
	time.Sleep(time.Second * 5)
	tellEndTurn()
	chanPomodoro <- true
}

func pomodoroBreak(chanBreak chan bool) {
	tellBeginSmallBreak()
	time.Sleep(time.Second * 5)
	tellEndSmallBreak()
	chanBreak <- true
}

func pomodoroLongBreak(chanLongBreak chan bool) {
	tellBeginLongBreak()
	time.Sleep(time.Second * 5)
	tellEndLongBreak()
	chanLongBreak <- true
}

func tellBeginTurn() {
	exec.Command("say", "Pomodoro round begins").Output()
}

func tellEndTurn() {
	exec.Command("say", "Round ended").Output()
}

func tellBeginSmallBreak() {
	exec.Command("say", "Have a small break!").Output()
}

func tellEndSmallBreak() {
	exec.Command("say", "This is the end of the small break. Let's go back to work!").Output()
}

func tellBeginLongBreak() {
	exec.Command("say", "Have a long break! You deserve it!").Output()
}

func tellEndLongBreak() {
	exec.Command("say", "This is the end of the long break. Let's go back to work!").Output()
}

func pomodoroService(chanPomodoro, chanBreak, chanLongBreak, chanDone chan bool) {
	fmt.Println("Pomodoro service started\n")
	for {
		select {

		case endTurn := <-chanPomodoro:
			_ = endTurn
			if currentTurn >= totalTurns {
				go pomodoroLongBreak(chanLongBreak)
				currentTurn = 1
			} else {
				currentTurn += 1
				go pomodoroBreak(chanBreak)
			}

		case endSmallBreak := <-chanBreak:
			_ = endSmallBreak
			go pomodoroTurn(chanPomodoro)

		case endLongBreak := <-chanLongBreak:
			_ = endLongBreak
			input := askAnotherSession()
			for input != "Y" && input != "N" && input != "y" && input != "n" {
				input = askAnotherSession()
			}
			if input == "Y" || input == "y" {
				go pomodoroTurn(chanPomodoro)
			} else {
				chanDone <- true
			}

		}
	}
}

func askAnotherSession() string {
	fmt.Println("Ready for another pomodoro session? (Y/N)")
	var input string
	fmt.Scanln(&input)
	return input
}
