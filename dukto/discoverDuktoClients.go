package dukto

import (
	"fmt"
	"net"
)

func udpBroadcastListen() {

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 4644,
	})

	if err != nil {
		fmt.Println("Error creating listener:", err)
		return
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading broadcast:", err)
			return
		}
		fmt.Println("Received broadcast message:", string(buf[:n]))
	}
}
