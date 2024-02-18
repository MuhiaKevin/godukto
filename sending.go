package main

import (
	"bytes"
	"encoding/binary"
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

	// get absolute path of folder i.e Music, Videos, etc instead of the full path
	filenameSlice := strings.Split(folderName, "/")
	rootName := filenameSlice[len(filenameSlice)-1]


	initialPacket, filesAndTheirPacket, err := createFolderInformation(folderName)


	// tcp conn send intial packet
	conn, err := net.Dial("tcp", "192.168.0.105:4644") 
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	// the information in the packet  will be  the folders total size, the number of files in the folder + 1 and the name of the folder
	_, err = conn.Write(initialPacket)
	if err != nil {
		log.Fatal(err)
	}

	// stream file
	// walkthrough the directory using the fullpath
	err = filepath.Walk(folderName, func(path string, info os.FileInfo, err error) error {
		// check file is not a directory
		if !info.IsDir() {
			// split the full path of the file into a slice of strings
			// objective is to get the path of the file relative to the root folder
			folderNamesSlice := strings.Split(path, "/") // separate files into slash

			// TODO: change this so that we don't need to check the index of the folder in the  each and every time
			posOfRootFolderInSlice := slices.Index(folderNamesSlice,rootName ) // get index of the root file from the file's full path

			// get each file's packet data that will be sentfirst to the dukto client before streaming the entire file
			newPath := strings.Join(folderNamesSlice[posOfRootFolderInSlice:], "/")
			// get the file's packet data
			filePack := filesAndTheirPacket[newPath]

			// open the file for reading
			file, err := os.Open(path)
			if err != nil {
				return err
			}


			// tcp send the file's packet first 
			_, err = conn.Write(filePack)
			if err != nil {
				log.Fatal(err)
			}


			// then using the same tcp connection stream the entire file.
			// the tcp connection will be reused to send all the files in the folder

			_, err = io.Copy(conn, file)
			if err != nil {
				return err
			}

			// close this file handler once the file has been sent
			file.Close()
		}

		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Folder sent Successfully!!!")

}

// structure of intial packet 
// number of bytes( 8 bytes) + total size of all files in the folder in little endian (8 bytes ) + folder/directory name +  A byte after name of folder (1 byte) + Ending part of the packet that is filled with 8 255 intergers ( 8 bytes)

func createFolderInformation(folderName string) ([]byte, map[string][]byte, error) {
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

// STRUCTURE OF FILE PACKET
// Path to file from the mainfolder + 1 byte + size of the file ( 8 bytes)


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


	filesAndtheirPackets[newPath] = filePacket.Bytes()
}
