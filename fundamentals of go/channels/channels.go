package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan string)

	go pinger(c)
	go printer(c)

	fmt.Println("Main is sleeping!")
	var input string
	fmt.Scanln((&input))

	fmt.Println("Main is done!")

	// Select Case
	Select()
}
