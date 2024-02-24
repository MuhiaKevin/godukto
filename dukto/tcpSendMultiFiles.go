package dukto

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)


func SendMultipleFiles(files []string, host string) error {
	client := fmt.Sprintf("%s:%d", host, port)
	totalSize := 0

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		fs, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		totalSize = totalSize + int(fs.Size())
		file.Close()
	}

	conn, err := net.Dial("tcp", client)
	if err != nil {
		return err
	}
	defer conn.Close()

	fisrtPack := CreateFirstPacket(totalSize, len(files))
	conn.Write(fisrtPack)

	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}

		fs, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}

		filenameSlice := strings.Split(filename, "/")
		correctName := filenameSlice[len(filenameSlice)-1]

		fileSize := fs.Size()

		filenamePack := CreateFilePacket(correctName, int(fileSize))

		conn.Write(filenamePack)

		io.Copy(conn, file)

		file.Close()
	}

	return nil
}
