package main

import (
	"fmt"
	"sync"
	"time"
)

func Even(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case num := <-ch:
			if num%2 == 0 {
				fmt.Printf("%d Even", num)
			} else {
				ch <- num
			}

		}
	}
}
func Odd(ch chan int, wg *sync.WaitGroup) {
	for {
		select {
		case num := <-ch:
			if num%2 != 0 {
				fmt.Printf("%d Odd", num)
			} else {
				ch <- num
			}

		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 2)
	go Odd(ch, &wg)
	go Even(ch, &wg)

	for i := 1; i <= 5; i++ {
		ch <- i

		time.Sleep(time.Millisecond * 1000)
	}
	close(ch)
	wg.Done()
	wg.Done()
	wg.Wait()
}
