/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package cmd

import (
	"log"

	"github.com/ddefrancesco/scopectl/handlers"
	"github.com/spf13/cobra"
)

// initCmd represents the initialize command
var initCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize the telescope",
	Long: `Initialize the telescope.

	Examples: scopectl initialize

	Usage: scopectl initialize
		  `,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println("initialize command called")
		pmap := make(map[string]string)

		pmap["address"], _ = cmd.Flags().GetString("location")
		scope_res, err := handlers.InitCommandHandler(pmap)

		if err != nil {
			return err
		}
		log.Printf("Init command responded: %s\n", scope_res.Response)
		return nil
	},
}

func init() {
	initCmd.Flags().StringP("location", "a", "", "Initialize scope. Set Time, Date. Convert via API street addresses into geographical coordinates (latitude and longitude). Set location via API.")
	rootCmd.AddCommand(initCmd)

}
