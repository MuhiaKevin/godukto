package networking

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"
)

func buildBroadCastMesage() []byte {
	header := make([]byte, 1)

	var messageBuf bytes.Buffer
	header[0] = 1
	message := "Muya at Muya (Golang)"

	messageBuf.Write(header)
	messageBuf.Write([]byte(message))

	return messageBuf.Bytes()
}

// func SendBroadcast(wg *sync.WaitGroup) {
func SendBroadcast() {
	message := buildBroadCastMesage()
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 4644,
	})
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	// defer wg.Done()
	for {
		_, err := conn.Write(message)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Sent broadcast message: ", string(message))
		time.Sleep(3 * time.Second)
	}
}
