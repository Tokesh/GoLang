package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generateData(link chan<- int) {
	defer close(link)
	for i := 0; i <= 5; i++ {
		link <- i
		time.Sleep(time.Second * 1)
	}

}
func showData(link <-chan int, done chan bool) {
	for b := range link {
		fmt.Println(b)
	}
	close(done)
}

func main() {
	buffer := make(chan int)
	done := make(chan bool)
	go generateData(buffer)
	go showData(buffer, done)
	<-done
}
