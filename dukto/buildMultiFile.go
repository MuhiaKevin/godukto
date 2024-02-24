package dukto

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"

	"net"
	"os"
	"strings"
)


func SendFiles(files []string, host string) error {
	client := fmt.Sprintf("%s:%d", host, port)

	if len(os.Args) < 2 {
		log.Fatal("Please enter more files")
	}

	files = os.Args[1:]
	totalSize := 0

	for _, file := range files {
		file, err := os.Open(file)
		if err != nil {
			return err
		}

		fs, err := file.Stat()
		if err != nil {
			return err
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
			return err
		}

		fs, err := file.Stat()
		if err != nil {
			return err
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

func CreateFirstPacket(totalSize, numOfFiles int) []byte {
	var firstPack bytes.Buffer

	numOfFilesBuf := make([]byte, 8)
	numOfFilesBuf[0] = byte(numOfFiles)
	totalSizeBuf := make([]byte, 8)

	binary.LittleEndian.PutUint64(totalSizeBuf, uint64(totalSize))

	firstPack.Write(numOfFilesBuf)
	firstPack.Write(totalSizeBuf)

	return firstPack.Bytes()
}

func CreateFilePacket(filename string, fileSize int) []byte {
	var filePack bytes.Buffer
	var filenameBuf bytes.Buffer

	nameFinishByte := make([]byte, 1)
	fileSizeBuf := make([]byte, 8)

	filenameBuf.Write([]byte(filename))
	binary.LittleEndian.PutUint32(fileSizeBuf , uint32(fileSize))

	filePack.Write(filenameBuf.Bytes())
	filePack.Write(nameFinishByte)
	filePack.Write(fileSizeBuf)

	return filePack.Bytes()
}
