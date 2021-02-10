package main

import (
	"fmt"
	"time"
)

func main() {
	//firstExample()
	ChannelsBetweenProcesses()
}



func ChannelsBetweenProcesses() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	go func() {
		for {
			test := <-channel
			fmt.Println(test)
		}
	}()
	var l string
	fmt.Scanln(&l)

}

func firstExample() {
	msg := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		msg <- "Hello World"
	}()

	result := <-msg
	fmt.Println(result)

}
