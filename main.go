package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// compare two lists of strings
func equal(a, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func monitor(hostname string) {
	addrs, _ := net.LookupHost(hostname)
	lookupTime := time.Now()

	log.Println(fmt.Sprintf("Watching %s for changes, initial result: %s", hostname, addrs))
	for {
		time.Sleep(1 * time.Second)
		newaddrs, _ := net.LookupHost(hostname)

		if !equal(addrs, newaddrs) {
			addrs = newaddrs
			log.Println(fmt.Sprintf("%s record changed to %s, %s since last change",
				hostname, newaddrs, time.Since(lookupTime)))
			lookupTime = time.Now()
		}
	}
}

func main() {
	// List of hostnames
	hostnames := os.Args[1:]
	if len(hostnames) == 0 {
		fmt.Println("hoststat monitors a list of hostnames for ip address changes")
		fmt.Println("Usage: hoststat gmail.com google.com")
		os.Exit(0)
	}
	for _, host := range hostnames {
		go monitor(host)
	}
	// hang forever
	select {}
}
