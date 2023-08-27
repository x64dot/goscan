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
		ip_target          string
		port_to_scan       string
		dummy_port_to_scan int
	)
	fmt.Println("Enter the ip address of the target you want to scan: ")
	fmt.Scanln(&ip_target)
	fmt.Println("Enter the port of the ip address you want to scan: ")
	fmt.Scanln(&dummy_port_to_scan)

	if dummy_port_to_scan > 65535 {
		fmt.Println("[!] Invalid port. You cannot specify a port higher than 65535!")
		os.Exit(0)
	}

	port_to_scan = strconv.Itoa(dummy_port_to_scan)
	time_out := time.Second * 3
	conn, err := net.DialTimeout("tcp", ip_target+":"+port_to_scan, time_out)

	if err != nil {
		fmt.Printf("[!] Port %v is not open!\n", port_to_scan)
		fmt.Println("\n\n ", err)
	} else {
		fmt.Printf("[+] Port %v is open!\n", port_to_scan)

	}

	defer conn.Close()

}
