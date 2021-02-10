package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go RunProcess("P1", 20)
	go RunProcess("P2", 20)
	wg.Wait()
}

func RunProcess(name string, total int)  {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	wg.Done()
}