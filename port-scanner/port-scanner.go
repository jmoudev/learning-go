package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func PortConnect(url string, port int) bool {
	address := fmt.Sprintf("%s:%d", url, port)
	conn, err := net.DialTimeout("tcp", address, time.Second * 2)

	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// TODO: Scan a range of ports concurrently with url input
// TODO: Add CLI
func main() {
	golang_url := "golang.org"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter port: ")
	input, _ := reader.ReadString('\n')
	port, _ := strconv.Atoi(strings.TrimSpace(input))

	fmt.Printf("Scanning port %d of %s\n", port, golang_url)

	ok := PortConnect(golang_url, port)
	if ok {
		fmt.Printf("Port %d OK\n", port)
	}
}
