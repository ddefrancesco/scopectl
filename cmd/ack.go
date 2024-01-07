/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ddefrancesco/scopectl/handlers"
	"github.com/spf13/cobra"
)

// alignCmd represents the align command
var ackCmd = &cobra.Command{
	Use:   "ack",
	Short: "Gets telescope alignment mode",
	Long: `Gets telescope Land, Polar and AltAz alignment modes.
	
	Examples: scopectl ack 
	
	Usage: scopectl ack`,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println("ack command is being called")

		scope_res, err := handlers.AckCommandHandler()
		if err != nil {
			fmt.Printf("error calling server API server: %s\n", err.Error())
			return err
		}
		fmt.Printf("ack command responded: %s\n", scope_res.Response)
		return nil
	},
}

func init() {

	rootCmd.AddCommand(ackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
