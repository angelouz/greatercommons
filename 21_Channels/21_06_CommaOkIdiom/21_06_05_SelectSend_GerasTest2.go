package main

import "fmt"

func send(c, quit chan int) {
	x := 3
	for {
		select {
		case c <- x:
			x += x
		case <-quit:
			fmt.Println("about to exit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {

		for {
			x := <-c
			fmt.Println(x)

			if x >= 10000 {
				quit <- 0
			}
		}
	}()
	send(c, quit)
}
