package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		ip_target    string
		port_to_scan string
	)
	fmt.Println("Enter the ip address of the target you want to scan: ")
	fmt.Scanln(&ip_target)
	fmt.Println("Enter the port of the ip address you want to scan: ")
	fmt.Scanln(&port_to_scan)

	dummy_port_to_scan, err := strconv.Atoi(port_to_scan)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}

	if dummy_port_to_scan < 1 || dummy_port_to_scan > 65535 {
		fmt.Println("[!] Invalid port. You cannot specify a port higher than 65535 or port 0!")
		os.Exit(0)
	}

	time_out := time.Second * 3
	conn, err := net.DialTimeout("tcp", ip_target+":"+port_to_scan, time_out)

	if err != nil {
		fmt.Printf("[!] Port %v is not open!\n", port_to_scan)
	} else {
		fmt.Printf("[+] Port %v is open!\n", port_to_scan)
		defer conn.Close()
	}
}
