package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

var result int
var m sync.Mutex

func main() {
	//WaitGroup()
	Mutex()
}

func WaitGroup()  {
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go RunProcess("P1", 20)
	go RunProcess("P2", 20)
	wg.Wait()
}

func Mutex()  {
	go RunProcess("P1", 20)
	go RunProcess("P2", 20)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result:", result)
}

func RunProcess(name string, total int)  {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		fmt.Println(name, "->", i, "Partial result:", result)
		m.Unlock()
	}
	//wg.Done()
}