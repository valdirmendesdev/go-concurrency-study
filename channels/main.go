package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//firstExample()
	//ChannelsBetweenProcesses()
	//ChannelsPlusWaitGroup()
	DeadlockChannel()
}

func DeadlockChannel() {
	channel := make(chan int)
	// Deadlock!
	//channel <- 10

	go func() {
		channel <- 10
	}()

	fmt.Println(<-channel)
}



func ChannelsPlusWaitGroup()  {
	var wg sync.WaitGroup
	channel := make(chan int)

	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}

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
