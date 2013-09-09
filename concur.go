package main

import (
	"fmt"
	"time"
)

//STARTPING OMIT
func Ping(ch chan *int) {
	for {
		i := <-ch
		time.Sleep(300 * time.Millisecond)
		*i++
		fmt.Println("Ping", *i)
		ch <- i
	}
}

func Pong(ch chan *int) {
	for {
		i := <-ch
		time.Sleep(300 * time.Millisecond)
		*i++
		fmt.Println("Pong", *i)
		ch <- i
	}
}

//STOPPONG OMIT

//STARTMAIN OMIT
func main() {
	ch := make(chan *int)
	go Ping(ch)
	go Pong(ch)
	var i int
	ch<-&i
	time.Sleep(5 * time.Second)

}

//STOPMAIN OMIT
