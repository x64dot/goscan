package main

import (
	"errors"
	"fmt"
	"log"
	"net"
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

	err := ValidateIP(ip_target)
	if err != nil {
		log.Fatalf("Invalid IP address: %v", err)
	}

	fmt.Println("Enter the port of the ip address you want to scan: ")
	fmt.Scanln(&port_to_scan)

	dummy_port_to_scan, err := strconv.Atoi(port_to_scan)

	if err != nil {
		log.Fatalf("Error converting port to integer: %v", err)
	}

	if dummy_port_to_scan < 1 || dummy_port_to_scan > 65535 {
		log.Fatal("[!] Invalid port. You cannot specify a port higher than 65535 or port 0!")
	}

	time_out := time.Second * 3
	conn, err := net.DialTimeout("tcp", ip_target+":"+port_to_scan, time_out)

	if err != nil {
		log.Printf("[!] Port %v is not open:", port_to_scan)
	} else {
		log.Printf("[+] Port %v is open!", port_to_scan)
		defer conn.Close()
	}
}

/* This function will simply check if the IP address provided is valid. */
/* This function will also work for IPv4 and IPv6 addresses.*/
func ValidateIP(ip string) error {
	ipt := net.ParseIP(ip)
	if ipt == nil {
		return errors.New("IP address is invalid")
	}
	return nil
}
