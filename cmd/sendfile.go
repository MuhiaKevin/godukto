/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "errors"
	// "fmt"
	"godukto/dukto"
	"log"
	"net"
	"os"

	// "godukto/dukto"
	// "log"

	"github.com/spf13/cobra"
)

// sendfileCmd represents the sendfile command
var sendfileCmd = &cobra.Command{
	Use:   "sendfile",
	Short: "Send file over lan",
	Long: `Send file over lan`,
	Run: start,
}

func start(cmd *cobra.Command, args []string) {
	// fmt.Println("sendfile called")

	// get filename
	file := args[0]

	// if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
	if _, err := os.Stat(file); err != nil {
		log.Fatal(err)
	} 

	// channel that gets dukto clients
	peers := make(chan net.IP)

	// discover other dukto apps
	go dukto.UdpBroadcastListen(peers)
	

	// read ip address from channel
	peerIP :=  <- peers
	// make sure message received is not bye so that it doesnt send a file to a closed tcp connection
	// log.Println("Received data from broadcat: ", peerIP.String())

	// sendfile
	// dukto.SendFile("./POTENTIAL_NEW_CONFIGS.zip", peerIP.String())
	dukto.SendFile(file, peerIP.String())
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
