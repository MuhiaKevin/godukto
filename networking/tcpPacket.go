package networking

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func GetCode(filename string) []byte {
	var full bytes.Buffer

	buf := make([]byte, 8)
	head := make([]byte, 1)

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	fileSize := fi.Size()
	binary.LittleEndian.PutUint32(buf, uint32(fileSize))

	full.Write(head)
	full.Write(buf)

	return full.Bytes()

}

func CreatePacketHeader(filename string) []byte {
	var header bytes.Buffer
	var full bytes.Buffer
	head := make([]byte, 7)

	filenameSlice := strings.Split(filename, "/")
	correctName := filenameSlice[len(filenameSlice)-1]
	fileNameBytes := []byte(correctName)

	code := GetCode(filename)

	head[0] = 1
	header.Write(head)
	header.Write(code)

	tail := code

	full.Write(header.Bytes())
	full.Write(fileNameBytes)
	full.Write(tail)

	return full.Bytes()

}

func main() {
	fileName := "/home/muhia/Documents/dotfiles-hyprland-main.zip"

	conn, err := net.Dial("tcp", "localhost:4644")
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}

	packet := CreatePacketHeader(fileName)
	// packet := getCode(fileName)

	conn.Write(packet)
}
