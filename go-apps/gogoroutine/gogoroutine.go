package main

import (
	"fmt"
	"time"
)

type MessageChannel = chan string

func greet(message string, c MessageChannel) {
	fmt.Printf("Normal %v\n", message)
}

func greetSlow(message string, c MessageChannel) {
	time.Sleep(1 * time.Second)
	fmt.Printf("Slow %s\n", message)
	// send data to the channel
	c <- message
	close(c)
}

func main() {
	// c := make(MessageChannel)

	// go greet("1", c)
	// go greet("2", c)
	// go greetSlow("111", c)
	// go greet("3", c)
	// go greet("3", c)
	// // wait for data come out from the channel
	// var message = <-c
	// // first message
	// if message != "" {
	// 	fmt.Printf("Last message: %v\n", message)
	// }

	// channels := make([]MessageChannel, 4)
	// for i, _ := range channels {
	// 	channels[i] = make(chan string)
	// 	go greetSlow(strconv.Itoa(i+1), channels[i])
	// }

	// for _, c := range channels {
	// 	<-c
	// }

	var channel MessageChannel = make(chan string)
	go greet("1", channel)
	go greet("2", channel)
	go greetSlow("3", channel)
	go greetSlow("4", channel)
	go greet("5", channel)
	go greet("6", channel)
	// for i := range 4 {
	// 	go greetSlow(strconv.Itoa(i+1), channel)
	// }

	for range <-channel {
	}
}
