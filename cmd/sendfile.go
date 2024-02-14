/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// sendfileCmd represents the sendfile command
var sendfileCmd = &cobra.Command{
	Use:   "sendfile",
	Short: "Send file over lan",
	Long: `Send file over lan`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sendfile called")
	},
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
