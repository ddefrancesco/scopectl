/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package cmd

import (
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
			log.Printf("error calling server API server: %s\n", err.Error())
			return err
		}
		log.Println("align ack command responded:")
		for _, v := range *scope_res {
			log.Printf("%s\n", v.Response)
		}
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
