package dukto

import (
	"fmt"
	"net"
)

func UdpBroadcastListen() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 4644,
	})

	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	defer conn.Close()
	fmt.Println("Listening for udp broadcast packets on port 4644")

	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading broadcast:", err)
			return
		}
		fmt.Printf("Received broadcast message: %v\n", string(buf[:n]))

		// write to channel

	}
}
