package main

import "fmt"

func hello(done chan bool) {
	fmt.Println("Hello world! goroutine")
	done <- true

}

func main() {
	var done chan bool
	if done == nil {
		fmt.Println("channel is a nil and I'm going to define it")
		done = make(chan bool)
		fmt.Printf("Type of a is %T\n", done)
	}
	go hello(done)
	<-done
	fmt.Println("main function")

}
