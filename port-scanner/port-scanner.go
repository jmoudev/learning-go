package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/urfave/cli/v3"
)

func PortConnect(wg *sync.WaitGroup, okPorts chan int, url string, port int) {
	defer wg.Done()

	address := net.JoinHostPort(url, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", address, time.Second*2)

	if err != nil {
		return
	}

	conn.Close()
	okPorts <- port
	return
}

func ScanPorts(url string, startPort int, endPort int) {
	numPorts := endPort - startPort

	var wg sync.WaitGroup
	wg.Add(numPorts)
	var okPorts = make(chan int)

	fmt.Printf("Scanning ports %s of %s\n", fmt.Sprintf("%d-%d", startPort, endPort), url)

	// Goroutine PortConnect asyncronously checks open ports
	for port := startPort; port < endPort; port++ {
		go PortConnect(&wg, okPorts, url, port)
	}

	go func() {
		wg.Wait()
		close(okPorts)
	}()

	for port := range okPorts {
		fmt.Printf("Port %d OK\n", port)
	}
}

func main() {
	cmd := &cli.Command{
		Name:      "port-scanner",
		Usage:     "Scan a number of ports for a given url",
		ArgsUsage: "<url>",
		Arguments: []cli.Argument{
			&cli.StringArg{
				Name: "url",
			},
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "startPort",
				Value: 1,
				Usage: "The lowest port in the port range",
			},
			&cli.IntFlag{
				Name:  "endPort",
				Value: 1000,
				Usage: "The highest port in the port range",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			url := cmd.StringArg("url")
			startPort := cmd.Int("startPort")
			endPort := cmd.Int("endPort")

			ScanPorts(url, startPort, endPort)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
