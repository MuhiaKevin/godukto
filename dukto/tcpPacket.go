package dukto

import (
	"bytes"
	"encoding/binary"
	"os"
	"strings"
)

func GetCode(filename string) ([]byte, error) {
	var codePacket bytes.Buffer

	buf := make([]byte, 8)
	head := make([]byte, 1)

	file, err := os.Open(filename)
	if err != nil {
		return []byte{}, err
	}

	fi, err := file.Stat()
	if err != nil {
		return []byte{}, err
	}

	fileSize := fi.Size()
	binary.LittleEndian.PutUint32(buf, uint32(fileSize))

	codePacket.Write(head)
	codePacket.Write(buf)

	return codePacket.Bytes(), nil
}

func CreatePacketHeader(filename string) ([]byte, error) {
	var header bytes.Buffer
	var fullHeaderPacket bytes.Buffer
	head := make([]byte, 7)

	filenameSlice := strings.Split(filename, "/")
	correctName := filenameSlice[len(filenameSlice)-1]
	fileNameBytes := []byte(correctName)

	code, err := GetCode(filename)
	if err != nil {
		return []byte{}, err
	}

	head[0] = 1
	header.Write(head)
	header.Write(code)

	tail := code

	fullHeaderPacket.Write(header.Bytes())
	fullHeaderPacket.Write(fileNameBytes)
	fullHeaderPacket.Write(tail)

	return fullHeaderPacket.Bytes(), nil
}
