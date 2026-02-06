package main

import (
	"fmt"
	"hello/greeting"
	"hello/names"
	"sync"
)

func PrintRandomGreeting(wg *sync.WaitGroup) {
	defer wg.Done()

	name := names.GetRandomName()
	message := greeting.Hello(name)
	fmt.Println(message)
}

func main() {
	numNames := 10

	var wgNames sync.WaitGroup
	var wgMessages sync.WaitGroup
	wgNames.Add(numNames)
	wgMessages.Add(numNames)

	var names_ = make(chan string)
	var messages = make(chan string)

	for range numNames {
		go func() {
			defer wgNames.Done()
			name := names.GetRandomNameDelayed()
			names_ <- name
		}()
	}

	go func() {
		wgNames.Wait()
		close(names_)
	}()

	for name := range names_ {
		go func() {
			defer wgMessages.Done()
			message := greeting.HelloDelayed(name)
			messages <- message
		}()
	}

	go func() {
		wgMessages.Wait()
		close(messages)
	}()

	for message := range messages {
		fmt.Println(message)
	}

}
