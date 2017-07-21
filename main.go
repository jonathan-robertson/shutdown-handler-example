package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go watchForShutdown(cancelFunc)

	messageChan := make(chan int)
	go receiver(messageChan)
	launchSenders(ctx, messageChan)

	log.Println("program terminated")
}

func watchForShutdown(cancelContext func()) {
	killswitch := make(chan os.Signal)
	signal.Notify(killswitch, os.Interrupt)
	for _ = range killswitch {
		fmt.Println(" interruption received!")
		cancelContext()
		close(killswitch)
	}
}

func receiver(c chan int) {
	for id := range c {
		log.Printf("received message from %d\n", id)
	}
}

func launchSenders(ctx context.Context, messageChan chan int) {
	var wg sync.WaitGroup
	numSenders := 3
	for i := 0; i < numSenders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sender(ctx, messageChan, numSenders, i)
		}()
		time.Sleep(1 * time.Second) // stagger startup of routines
	}

	wg.Wait()
	close(messageChan) // tell receiver that its work is complete
}

func sender(ctx context.Context, c chan int, numSenders int, id int) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d noticed termination signal and shut down\n", id)
			return
		case c <- id:
			time.Sleep(time.Duration(numSenders) * time.Second)
		}
	}
}
