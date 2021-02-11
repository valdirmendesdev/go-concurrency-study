package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := funnel(generate("hello go"), generate("Hello World!"))
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Finished!")
}

func generate(s string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("String: %s - Value: %d", s, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()
	return channel
}

func funnel(ch1, ch2 <-chan string) <-chan string {
	channel := make(chan string)
	go func() {
		for  {
			channel <- <-ch1
		}
	}()
	go func() {
		for  {
			channel <- <-ch2
		}
	}()
	return channel
}
