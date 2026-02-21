package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func PortConnect(wg *sync.WaitGroup, ok_ports chan int, url string, port int) {
	defer wg.Done()

	address := net.JoinHostPort(url, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, time.Second*2)

	if err != nil {
		return
	}

	conn.Close()
	ok_ports <- port
	return
}

// TODO: Add CLI
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter url: ")
	input, _ := reader.ReadString('\n')
	url := strings.TrimSpace(input)

	port_start := 10
	port_end := 1000
	num_ports := port_end - port_start

	var wg sync.WaitGroup
	wg.Add(num_ports)
	var ok_ports = make(chan int)

	fmt.Printf("Scanning ports %s of %s\n", fmt.Sprintf("%d-%d", port_start, port_end), url)

	// Goroutine PortConnect asyncronously checks open ports
	for port := port_start; port < port_end; port++ {
		go PortConnect(&wg, ok_ports, url, port)
	}

	go func() {
		wg.Wait()
		close(ok_ports)
	}()

	for port := range ok_ports {
		fmt.Printf("Port %d OK\n", port)
	}
}
