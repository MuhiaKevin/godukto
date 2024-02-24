package dukto

import (
	"fmt"
	"io"
	"net"
	"os"
)

const port = 4644

func SendFile(fileName string, host string) error {
	client := fmt.Sprintf("%s:%d", host, port)

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	conn, err := net.Dial("tcp", client)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet, err := CreatePacketHeader(fileName)
	if err != nil {
		return err
	}

	conn.Write(packet)

	fmt.Println("Sending file ", fileName)

	// https://itnext.io/optimizing-large-file-transfers-in-linux-with-go-an-exploration-of-tcp-and-syscall-ebe1b93fb72f
	_, err = io.Copy(conn, file)
	if err != nil {
		return err
	}

	fmt.Println("File sent successfully")
	return nil
}

func ReceiveFile() {

}
