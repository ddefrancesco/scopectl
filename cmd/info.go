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
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Gets telescope informations",
	Long: `Gets  informations about telescope functional parameters.
	
	Examples: scopectl info --infos=altitude 
	          scopectl info --infos=azimuth,declination
			  scopectl info --infos=all
	
	Usage: scopectl scopectl info --infos=<parameters>`,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println("info command is being called")
		infos, err := cmd.Flags().GetString("infos")
		if err != nil {
			fmt.Printf("error retrieving infos: %s\n", err.Error())
			return err
		}
		params := make(map[string]string)
		params["info"] = infos
		scope_res, err := handlers.InfoCommandHandler(params)
		if err != nil {
			fmt.Printf("error calling server API server: %s\n", err.Error())
			return err
		}
		fmt.Println("info command responded:")
		for _, v := range *scope_res {
			fmt.Printf("%s\n", v.Response)
		}
		return nil
	},
}

func init() {
	infoCmd.Flags().StringP("infos", "i", "", "informations query parameters")
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
