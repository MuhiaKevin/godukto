package dukto

import (
	"fmt"
	"io"
	"net"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SendFile(fileName string) {
	file, err := os.Open(fileName)

	CheckErr(err)

	defer file.Close()

	conn, err := net.Dial("tcp", "192.168.1.195:4644")
	defer conn.Close()

	CheckErr(err)

	packet := CreatePacketHeader(fileName)

	conn.Write(packet)

	_, err = io.Copy(conn, file)
	CheckErr(err)

	fmt.Println("File sent successfully")
}

func ReceiveFile() {

}
