package main

import (
	"fmt"
	"hello/greeting"
)

func main() {
	var name string
	fmt.Print("Type name: ")
  	fmt.Scan(&name)
  
	message := greeting.Hello(name)
	fmt.Println(message)
}