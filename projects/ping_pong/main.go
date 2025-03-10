package main

import (
	"fmt"
	"time"
)

func ping(c chan string){
	for i:=0; ; i++{
	c <- "ping"
	}
}

func pong(c chan string){
	for i:=0; ; i++{
	c <- "pong"
	}
}

func show(c chan string){
	for{
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main(){
	var c chan string = make(chan string)

	go ping(c)
	go show(c)
	go pong(c)

	var input string
	fmt.Scanln(&input)
}

