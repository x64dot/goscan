package main

import (
	"errors"
	"flag"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var generalTimeout = time.Second * 3

func main() {
	var (
		targetIP  string
		startPort string
		endPort   string
		scanAll   bool 
	)

	flag.StringVar(&targetIP, "ip", "", "IP address to scan")
	flag.StringVar(&startPort, "p", "", "Start port to scan from | scan one port")
	flag.StringVar(&endPort, "pp", "", "End port to stop scan from")
	flag.BoolVar(&scanAll, "ap", false, "Scan all ports") 

	flag.Parse()
	if targetIP == "" {
		log.Fatal("[!] Please enter an IP address.")
	}

	err := ValidateIP(targetIP)
	if err != nil {
		log.Fatalf("[!] %v", err)
	}

	if len(startPort) > 0 && endPort == "" {
		err := ScanOnePort(targetIP, startPort)

		if err == nil {
			log.Printf("[+] Port: %v is open!", startPort)
		} else {
			log.Printf("[-] Port: %v is closed!", startPort)
		}
	}

	if len(startPort) > 0 && len(endPort) > 0 {
		err := ScanPorts(targetIP, startPort, endPort)
		if err != nil {
			log.Fatalf("[!] Error scanning ports: %v", err)
		}
	}

	if scanAll {
		ScanAll(targetIP)
	}
}

func ValidateIP(ip string) error {
	ipt := net.ParseIP(ip)
	if ipt == nil {
		return errors.New("[!] IP address is invalid")
	}
	return nil
}

func ScanOnePort(ip string, port string) error {
	conn, err := net.DialTimeout("tcp", ip+":"+port, generalTimeout)

	if err == nil {
		conn.Close()
		return nil
	} else {
		return err
	}
}

func ScanPorts(ip string, sport string, eport string) error {
	startPort, err := strconv.Atoi(sport)
	if err != nil {
		log.Print()
	}
	endPort, err := strconv.Atoi(eport)
	if err != nil {
		log.Print()
	}

	var wg sync.WaitGroup

	for ; startPort <= endPort; startPort++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			target := ip + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout("tcp", target, generalTimeout)

			if err == nil {
				conn.Close()
				log.Printf("[+] %v Port is open!\n", port)
			}
		}(startPort)
	}

	wg.Wait()
	return nil
}

func ScanAll(ip string) {
	var wg sync.WaitGroup

	for port := 1; port < 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			target := ip + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout("tcp", target, generalTimeout)

			if err == nil {
				conn.Close()
				log.Printf("[+] %v Port is open!\n", port)
			}

		}(port)

	}
	wg.Wait()
}
