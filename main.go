package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func shutdownWaiter(s chan os.Signal, numWorkers int) {
	signal.Notify(s, os.Interrupt)
	go func() {
		for _ = range s {
			// sig is a ^C, handle it
			fmt.Println(" interruption received!")
			close(s)
		}
	}()
}

func printWorker(c chan time.Time) {
	for myTime := range c {
		fmt.Println(myTime)
	}
}

func loadWorker(c chan time.Time, s chan os.Signal, numWorkers int, num int) {
	for {
		select {
		case <-s:
			fmt.Println("Worker", num, "saw stop signal")
			return
		case c <- time.Now():
			time.Sleep(time.Duration(numWorkers) * time.Second)
		}
	}
}

func main() {
	c := make(chan time.Time)
	shutdownChan := make(chan os.Signal, 1)
	var wg sync.WaitGroup
	numWorkers := 3

	go shutdownWaiter(shutdownChan, numWorkers)

	go printWorker(c)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		time.Sleep(1 * time.Second)
		go func(num int) {
			defer wg.Done()
			loadWorker(c, shutdownChan, numWorkers, num)
		}(i)

	}

	wg.Wait() // Wait for loadWorkers to reach Done
	close(c)  // Tell printWorker to not expect any new entries to chan
	fmt.Println("Program terminated")
}
