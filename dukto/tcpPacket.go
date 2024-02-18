package dukto

import (
	"bytes"
	"encoding/binary"
	"os"
	"strings"
)

// structure 
// 1 byte (value is zero) + size of the file(8 bytes)

func GetCode(filename string) ([]byte, error) {
	// this will store information about file size
	var codePacket bytes.Buffer

	// this will save the size of the file in bytes. Size must be 8 bytes
	buf := make([]byte, 8)

	// This is the extra byte that needs to be appended. This value is zero
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
	// store the size of the file 
	binary.LittleEndian.PutUint32(buf, uint32(fileSize))

	// add everything together
	codePacket.Write(head)
	codePacket.Write(buf)

	return codePacket.Bytes(), nil
}

func CreatePacketHeader(filename string) ([]byte, error) {
	var header bytes.Buffer
	var fullHeaderPacket bytes.Buffer
	head := make([]byte, 7)

	// get exact name of file instead of the path
	filenameSlice := strings.Split(filename, "/")
	correctName := filenameSlice[len(filenameSlice)-1]
	fileNameBytes := []byte(correctName)

	// get the first section of the packet
	code, err := GetCode(filename)
	if err != nil {
		return []byte{}, err
	}

	head[0] = 1
	header.Write(head)
	header.Write(code)

	tail := code

	// first 16 bytes includes 8 bytes but the first byte is 
	// second 8 bytes include the size of the file
	fullHeaderPacket.Write(header.Bytes()) 
	// next part includes the name of the file
	fullHeaderPacket.Write(fileNameBytes)
	// the last past is the bytes that information on the size of the file
	fullHeaderPacket.Write(tail)

	// to get filename from bytes
	// full := fullHeaderPacket.Bytes()	
	// fmt.Println(string(full[16 : (16 + len(filename))]))


	return fullHeaderPacket.Bytes(), nil
}
