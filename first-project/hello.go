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

	var wg sync.WaitGroup
	wg.Add(numNames)

	for range numNames {
		go PrintRandomGreeting(&wg)
	}
	wg.Wait()
}
