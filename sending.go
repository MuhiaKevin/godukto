package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func main() {
	folderName := "/home/muhia/Downloads/Videos" // make sure the folder doesnt end with "/"


	// get absolute path of folder
	filenameSlice := strings.Split(folderName, "/")
	rootName := filenameSlice[len(filenameSlice)-1]


	initialPacket, filesAndTheirPacket, err := createInitialPacket(folderName)
	// _, filesAndTheirPacket, err := createInitialPacket(folderName)


	// tcp conn send intial packet
	conn, err := net.Dial("tcp", "192.168.0.105:4644") 
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_, err = conn.Write(initialPacket)
	if err != nil {
		log.Fatal(err)
	}

	// stream file
	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			filenameSlice := strings.Split(path, "/") // separate files into slash
			folderNamePost := slices.Index(filenameSlice,rootName ) 

			// get the files packet
			newPath := strings.Join(filenameSlice[folderNamePost:], "/")
			filePack := filesAndTheirPacket[newPath]

			fmt.Println(filePack)

			// open the file 
			file, err := os.Open(path)
			if err != nil {
				return err
			}


			// tcp send the file's packet
			_, err = conn.Write(filePack)
			if err != nil {
				log.Fatal(err)
			}


			// stream the tcp packet

			_, err = io.Copy(conn, file)
			if err != nil {
				return err
			}

			file.Close()
		}

		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(filesAndTheirPacket)

}

// structure of intial packet 
// number of bytes( 8 bytes) + total size of all files in the folder in little endian (8 bytes ) + folder/directory name +  A byte after name of folder (1 byte) + Ending part of the packet that is filled with 8 255 intergers ( 8 bytes)

func createInitialPacket(folderName string) ([]byte, map[string][]byte, error) {
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
	numOfFiles, folderSize, filesAndTheirPacket,  err := getFolderInfo(folderName, rootName)
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

func sendData(data []byte)  {
	conn, err := net.Dial("tcp", "127.0.0.1:9090") 
	if err != nil {
		log.Fatal(err)
	}
	
	defer conn.Close()
	
	_, err = conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	
}

func getFolderInfo(folderName, rootFolderName string) (int64, int64, map[string][]byte, error) {
	var totalSize int64 // get total size of all the files in the folder
	var numOfFiles int64 // get the total number of all the files in the folder 
	filesAndTheirPacket := make(map[string][]byte) // associate each file with the intial section of packet when sending the files

	err := filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			getFileFromFolderName(rootFolderName, path, info.Size(), filesAndTheirPacket )
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


func getFileFromFolderName(folderName, path string, fileSize int64, filesAndtheirPackets map[string][]byte )  {
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

	// fmt.Println(filePacket.Bytes())

	filesAndtheirPackets[newPath] = filePacket.Bytes()

	// sendData(filePacket.Bytes())
}
