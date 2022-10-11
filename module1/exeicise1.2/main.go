package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int, 10)
	go func() {
		tickerin := time.NewTicker(1 * time.Second)
		for _ = range tickerin.C {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("putting from child thread: ", n)
			ch <- n
		}
	}()

	time.Sleep(15 * time.Second)

	tickerout := time.NewTicker(1 * time.Second)
	for _ = range tickerout.C {
		fmt.Println("receiving from father thread: ", <-ch)
	}
}
