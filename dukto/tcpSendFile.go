package dukto

import (
	"fmt"
	"io"
	"net"
	"os"
)

const port = 4644

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func SendFile(fileName string, host string) {
	client := fmt.Sprintf("%s:%d", host, port)

	file, err := os.Open(fileName)

	CheckErr(err)

	defer file.Close()

	conn, err := net.Dial("tcp", client)
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
