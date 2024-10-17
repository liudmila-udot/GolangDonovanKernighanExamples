package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

func main() {

	// Make channels to pass fatal errors in WaitGroup
	fatalErrors := make(chan error)
	wgDone := make(chan bool)
	go func() {
		fatalErrors <- errors.New("MIla")
	}()
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		log.Println("Waitgroup 1")
		// Do Something...
		wg.Done()
	}()
	go func() {
		log.Println("Waitgroup 2")
		// Example function which returns an error
		err := ReturnsError()
		if err != nil {
			fatalErrors <- err
		}
		fmt.Println()
		wg.Done()
		//time.Sleep()
	}()

	// Important final goroutine to wait until WaitGroup is done
	go func() {
		wg.Wait()
		close(wgDone)
	}()

	// Wait until either WaitGroup is done or an error is received through the channel
	select {
	case <-wgDone:
		// carry on
		break
	case err := <-fatalErrors:
		close(fatalErrors)
		log.Fatal("Error: ", err)
	}

	log.Println("Program executed successfully")
}

func ReturnsError() error {
	return errors.New("Example error on golangcode.com")
}
