package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	group   sync.WaitGroup
	channel chan string
)

func Forword() {
	for i := 0; i <= 10; i++ {
		channel <- fmt.Sprintf("Forword i=%v", i) // send message
		time.Sleep(time.Millisecond)
	}

	fmt.Println("Forword finish its task ")
	group.Done()
}

//Printng both the values
func Reverse() {
	for i := 10; i >= 0; i-- {
		if s, ok := <-channel; ok { // receive message
			fmt.Println(s) //Print the value of Forword

			fmt.Printf("   Reverse i=%v\n", i) // Print the value of Reverse
			time.Sleep(time.Millisecond)
		}
	}

	fmt.Println(" Reverse finished task")
	group.Done()
}

func main2() {
	channel = make(chan string)
	group.Add(2)
	go Forword()
	go Reverse()
	group.Wait()
	close(channel) // close it so nothing is waiting
	for s := range channel {
		fmt.Println(s)
	}
}
