package dukto

import (
	"fmt"
	"net"
	"strings"
	// "strings"
)

func UdpBroadcastListen(peers chan net.IP) {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 4644,
	})

	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	defer conn.Close()
	// fmt.Println("Listening for udp broadcast packets on port 4644")
	fmt.Println("Waiting for other dukto applications...")

	for {
		buf := make([]byte, 1024)
		n, udpAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading broadcast:", err)
			return
		}

		message := string(buf[:n])

		if strings.Contains(message, "Bye Bye") {
			fmt.Printf("Device with IP %v is saying %s",  udpAddr.String(), message)
		} else {
			// fmt.Printf("Received broadcast message: %v with ip address: %v\n", string(buf[:n]), udpAddr.String())
			fmt.Printf("Found  device %v\n", message)
			// fmt.Printf("Device sent Message: %v\n", udpAddr.String())
			// fmt.Printf("Received broadcast message: %v with ip address: %v\n", message, udpAddr.String())
			// write to channel

			peers <- udpAddr.IP
		}
	}
}
