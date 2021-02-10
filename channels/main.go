package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		msg <- "Hello World"
	}()

	result := <-msg
	fmt.Println(result)
}
