package main

import (
	"flag"
	"fmt"
	"net"
	"sort"
	"time"
)

type Silent struct{}

// defining a worker, taking two channels of int ports and results as params
func worker(ports, results chan int, target string) {
	for p := range ports { // for loop to iterate through the ports channel until closed.
		address := fmt.Sprintf("%s:%d", target, p) // generates string representation of the remote host and the port being scanned. 'p' is the current port
		conn, err := net.Dial("tcp", address)      // with the given address, attempt to open a TCP conneection.
		if err != nil {
			results <- 0 // check to see if the conn is successful. if there is an error, close with a value of 0 which is sent to the results channel.
			continue     // skip to the next iteration of the loop
		}
		conn.Close() // when we do have a successful connection, we close it
		results <- p // also when we are successful, send the value of p (current port being scanned) to the results
		fmt.Printf("Port %d scanned\n", p)
	}
}

func main() {
	startTime := time.Now()
	urlArg := flag.String("url", "", "source url")
	flag.Parse()

	// initalize ports and results and empty slice called openports
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	if *urlArg == "" {
		fmt.Println("Please specify a target URL using the --url flag")
		return
	}
	target := *urlArg

	// creates worker goroutines for scanning ports. cap(ports) returns the capacity of the ports (100), goroutine reads from the ports and writes to results channel
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, target)
	}

	// go routine that sends each port number in the given range to the ports channel
	// @goroutine - ltes us scan multiple ports simultaenously. in this case, we can scan 100 ports concurrently
	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
			// fmt.Printf("Sent port %d to worker\n", i)
		}
	}()

	// we then collect the results from the results channel, and appends the open ports to the slice defined

	for port := range results {
		if port != 0 {
			openports = append(openports, port)
			fmt.Printf("Port %d is open\n", port)
		} else {
			// fmt.Printf("Port %d is closed\n", port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	fmt.Println("Open ports:")
	for _, port := range openports { // finallt we want to sort and print the port numbers
		fmt.Printf("%d open\n", port)
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsedTime)
}
