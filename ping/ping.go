package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"
	"fmt"
)

func foo(channel chan string) {
	for {
		// TODO: Write an infinite loop of sending "pings" and receiving "pongs"
		channel <- "ping"
		fmt.Println("Foo is sending: ping")
		message := <- channel
		fmt.Println("Foo is receiving:", message)
	}
}

func bar(channel chan string) {
	for {
		// TODO: Write an infinite loop of receiving "pings" and sending "pongs"
		message := <- channel
		fmt.Println("Bar is receiving:", message)
		channel <- "pong"
		fmt.Println("Bar is sending: pong")
	}
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	channel := make(chan string)
	go foo(channel) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(channel)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
