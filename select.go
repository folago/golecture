package main

import (
	"fmt"
	"math/rand"
	"time"
)

//START OMIT
func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(3)))
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(3)))
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// STOP OMIT
