package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var count int

func function1() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Function 1 =>", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	waitGroup.Done()
}

func function2() {
	for i := 0; i <= 5; i++ {
		fmt.Println("Function 2 =>", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	waitGroup.Done()
}

func main() {
	waitGroup.Add(2)
	go function1()
	go function2()
	waitGroup.Wait()
	fmt.Println("Print Count: ", count)
}
