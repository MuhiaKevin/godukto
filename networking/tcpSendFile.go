package networking

import (
	"fmt"
	"io"
	"net"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
}

// func SendFile(wg *sync.WaitGroup) {
func SendFile() {

	fileName := "/home/muhia/Documents/dotfiles-hyprland-main.zip"
	// Connect to the server
	conn, err := net.Dial("tcp", "192.168.1.195:4644")
	// conn, err := net.Dial("tcp", "localhost:4644")
	checkErr(err)

	defer conn.Close()
	// Open the file
	file, err := os.Open("/home/muhia/Documents/dotfiles-hyprland-main.zip")

	checkErr(err)

	defer file.Close()

	packet := CreatePacketHeader(fileName)

	conn.Write(packet)

	// Read the contents of the file into a buffer
	_, err = io.Copy(conn, file)
	checkErr(err)

	fmt.Println("File sent successfully")

	// wg.Done()
}
