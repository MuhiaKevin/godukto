package cmd

import (
	"fmt"
	"godukto/dukto"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// sendfolderCmd represents the sendfolder command
var sendfolderCmd = &cobra.Command{
	Use:   "sendfolder",
	Short: "Go dukto send a folder",
	Long: "Go dukto send a folder. Make sure folder doesn't end with a black slash",
	Run: startSendFolder,
}



func startSendFolder(cmd *cobra.Command, args []string) {
	// chcek that the file has been set
	if len(args) == 0 || len(args) > 1{
		log.Fatal("ERROR: set a folder  to send")
	}

	// TODO: send udp broadcast message to tell dukto clients you are there

	// list of clients that have been detected
	duktoClientsSeverd := make(map[string]string)

	// get filename
	folder := strings.TrimSuffix(args[0], "/")

	// check if file actually exists
	// if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
	if _, err := os.Stat(folder); err != nil {
		log.Fatal(err)
	} 


	// TODO: Get inital packet first while waiting for a dukto client to be discovered
	initialPacket, filesAndTheirPacket, err := dukto.CreateFolderInformation(folder)
	if  err != nil {
		log.Fatal(err)
	} 


	// channel that gets dukto clients from broadcast
	peers := make(chan dukto.DuktoClient)

	// discover other dukto apps and write to the peer channel
	go dukto.UdpBroadcastListen(peers)
	

	// read ip address from channel
	// peerIP :=  <- peers
	// // make sure message received is not bye so that it doesnt send a file to a closed tcp connection
	// // log.Println("Received data from broadcat: ", peerIP.String())
	//
	// // sendfile
	// // dukto.SendFile("./POTENTIAL_NEW_CONFIGS.zip", peerIP.String())
	// dukto.SendFile(file, peerIP.String())


	for {
		// message from udpBroadcast
		duktoClient, ok := <- peers

		// check if channel is closed 
		if ok == false {
			log.Println("Received data from broadcat: ", duktoClient.IP)
		} else {
			// have a list of dukto clients you have already sent them the file 
			// if you have already sent them the file then dont send the file to them again
			// else send it to them

			// check if the dukto client is in the list
			if v, ok := duktoClientsSeverd[duktoClient.IP]; ok {
				fmt.Printf("%v ALready Exists\n", v)
			} else {
				// if not then start a goroutine and  send the file to the dukto client
				go dukto.SendFolder(folder, duktoClient.IP, initialPacket, filesAndTheirPacket)
			}
			

			// if ipAddr, ok := duktoClientsSeverd[duktoClient.Name]; ok {
			// 	go dukto.SendFile(file, ipAddr)
			// }

			// add the dukto client to the list of served dukto clients
			duktoClientsSeverd[duktoClient.IP] = duktoClient.Name

			// fmt.Println(duktoClientsSeverd)
			// fmt.Println("Clients available", len(duktoClientsSeverd))
		}
	}
}

func init() {
	rootCmd.AddCommand(sendfolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendfolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendfolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
