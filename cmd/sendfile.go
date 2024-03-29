/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"godukto/dukto"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// sendfileCmd represents the sendfile command
var sendfileCmd = &cobra.Command{
	Use:   "sendfiles",
	Short: "Send one or more files to a dukto client or to multiple dukto clients",
	Long: `Send one file or multiple files to one or more dukto clients`,
	Run: startSendFile,
}

func startSendFile(cmd *cobra.Command, args []string) {
	// chcek that the file has been set
	if len(args) == 0 {
		log.Fatal("ERROR: set a single file  to send")
	}

	// list of clients that have been detected
	duktoClientsSeverd := make(map[string]string)

	// perdiocally send udp broadcast message to make other dukto clients aware of you
	// err := dukto.SendUdpBroadcast()
	// if err != nil {
	// 	log.Fatal(err)
	// } 

	// perdiocally send udp broadcast message to make other dukto clients aware of you
	// TODO: Find out why sometimes not working
	go dukto.SendUdpBroadcast()

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
			if _, ok := duktoClientsSeverd[duktoClient.IP]; !ok {
				// if not then start a goroutine and  send the file to the dukto client
				// send one or many files
				if len(args) > 1 {
					files := args[1:] 

					go dukto.SendMultipleFiles(files, duktoClient.IP)
				} else {
					file := args[0]

					if _, err := os.Stat(file); err != nil {
						log.Fatal(err)
					} 

					go dukto.SendFile(file, duktoClient.IP)
				}
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
	rootCmd.AddCommand(sendfileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendfileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendfileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
