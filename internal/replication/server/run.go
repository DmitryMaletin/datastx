package main

import (
	"fmt"
	"time"
)

func service1() {
	for {
		fmt.Println("Running service 1")
		time.Sleep(time.Second)
	}
}

func service2() {
	for {
		fmt.Println("Running service 2")
		time.Sleep(time.Second)
	}
}

func runIt() {
	// Start service 1 in a new goroutine
	go service1()

	// Start service 2 in a new goroutine
	go service2()

	// Wait indefinitely
	select {}
}
