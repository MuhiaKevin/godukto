package dukto

import (
	"fmt"
	"net"
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
	fmt.Println("Listening for udp broadcast packets on port 4644")

	for {
		buf := make([]byte, 1024)
		n, udpAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("Error reading broadcast:", err)
			return
		}

		fmt.Printf("Received broadcast message: %v with ip address: %v\n", string(buf[:n]), udpAddr.String())
		// write to channel

		peers <- udpAddr.IP
	}
}
