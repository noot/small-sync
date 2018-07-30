package main

import (
	"fmt"
)

func input(messages chan string, done chan bool) {
	done_out := make(chan bool)
	var input string

	go output(messages, done_out)
	fmt.Scanln(&input)
	messages <- input

	<-done_out

	go output(messages, done_out)
	fmt.Scanln(&input)
	messages <- input

	<-done_out

	done <- true
}

func output(messages chan string, done chan bool) {
	msg := <-messages
	fmt.Println(msg)

	done <- true
}

func main() {
	fmt.Println("starting execution...")
	messages := make(chan string, 2)
	done := make(chan bool)

	go input(messages, done)

	<-done
}