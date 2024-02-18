package dukto

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"slices"
	"strings"
)


func CreateFolderInformation(folderName string) ([]byte, map[string][]byte, error) {
	// will be used to concatenate all the sections of the packet
	var intialPacket bytes.Buffer

	// get absolute path of folder
	filenameSlice := strings.Split(folderName, "/")
	rootName := filenameSlice[len(filenameSlice)-1]
	rootNameBytes := []byte(rootName)


	// set bytes that will be used to save the information
	numOfFoldersByteSlice := make([]byte, 8)
	totalSizeBytes := make([]byte, 8)
	afterFileNameBytes := make([]byte, 1)
    endingBytes := bytes.Repeat([]byte{255}, 8)


	// get the total size of the folder and how many files it has
	numOfFiles, folderSize, filesAndTheirPacket,  err := GetFolderInfo(folderName, rootName)
	if err != nil {
		return []byte{}, map[string][]byte{},  err
	}

	// fmt.Printf("%d bytes\n", folderSize)
	// fmt.Printf("There are %d files in the folder\n", numberOfFiles)

	// set the number of files in the number of folders bytes slice
	binary.LittleEndian.PutUint64(numOfFoldersByteSlice, uint64(numOfFiles))
	// set the total size of files in the folders inside the bytes slice
	binary.LittleEndian.PutUint64(totalSizeBytes, uint64(folderSize))


	// add all the sections into one single byte
	intialPacket.Write(numOfFoldersByteSlice)
	intialPacket.Write(totalSizeBytes)
	intialPacket.Write(rootNameBytes)
	intialPacket.Write(afterFileNameBytes)
	intialPacket.Write(endingBytes)

	return intialPacket.Bytes(), filesAndTheirPacket, nil
}

func GetFolderInfo(folderName, rootFolderName string) (int64, int64, map[string][]byte, error) {
	var totalSize int64 // get total size of all the files in the folder
	var numOfFiles int64 // get the total number of all the files in the folder 
	filesAndTheirPacket := make(map[string][]byte) // associate each file with the intial section of packet when sending the files

	err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			GetFileFromFolderName(rootFolderName, path, info.Size(), filesAndTheirPacket )
			totalSize += info.Size()
		}

		numOfFiles += 1
		return err
	})

	if err != nil {
		return 0, 0, map[string][]byte{} , err
	}


	return numOfFiles,totalSize,filesAndTheirPacket , err
}

// STRUCTURE OF FILE PACKET
// Path to file from the mainfolder + 1 byte + size of the file ( 8 bytes)


func GetFileFromFolderName(folderName, path string, fileSize int64, filesAndtheirPackets map[string][]byte )  {
	filenameSlice := strings.Split(path, "/") // separate files into slash
	folderNamePost := slices.Index(filenameSlice, folderName) // get position of root folder given the path

	newPath := strings.Join(filenameSlice[folderNamePost:], "/") // get  change this to some place
	// fmt.Println(newPath)


	var filePacket bytes.Buffer
	filePacket.Write([]byte(newPath))
	afterFileNameBytes := make([]byte, 1)
	filePacket.Write(afterFileNameBytes)

	totalSizeBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(totalSizeBytes, uint64(fileSize))
	filePacket.Write(totalSizeBytes)


	filesAndtheirPackets[newPath] = filePacket.Bytes()
}
