package dukto

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func buildUdpBroadCastMesage() []byte {
	header := make([]byte, 1)

	var messageBuf bytes.Buffer
	header[0] = 1
	message := "Chifu wa Kizunu (Golang)"

	messageBuf.Write(header)
	messageBuf.Write([]byte(message))

	return messageBuf.Bytes()
}

func SendUdpBroadcast() {
	message := buildUdpBroadCastMesage()
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 4644,
	})
	defer conn.Close()

	CheckErr(err)

	for {
		_, err := conn.Write(message)
		CheckErr(err)

		fmt.Println("Sent broadcast message: ", string(message))
		time.Sleep(3 * time.Second)
	}
}
