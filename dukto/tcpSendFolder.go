package dukto

import (
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"fmt"
)


func SendFolder(folderName string, host string) error {
	// folderName := "/home/muhia/Downloads/Videos" // make sure the folder doesnt end with "/"
	client := fmt.Sprintf("%s:%d", host, port)
	fmt.Println(folderName)

	// get absolute path of folder i.e Music, Videos, etc instead of the full path
	filenameSlice := strings.Split(folderName, "/")
	rootName := filenameSlice[len(filenameSlice)-1]


	// TODO: Get inital packet first while waiting for a dukto client to be discovered
	initialPacket, filesAndTheirPacket, err := CreateFolderInformation(folderName)

	// tcp conn send intial packet
	conn, err := net.Dial("tcp", client) 
	if err != nil {
		return err
	}

	defer conn.Close()

	// the information in the packet  will be  the folders total size, the number of files in the folder + 1 and the name of the folder
	_, err = conn.Write(initialPacket)
	if err != nil {
		return err
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
				return err
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
		return err
	}

	log.Println("Folder sent Successfully!!!")
	return nil
}
