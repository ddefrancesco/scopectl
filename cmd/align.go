/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/ddefrancesco/scopectl/handlers"
	"github.com/spf13/cobra"
)

// alignCmd represents the align command
var alignCmd = &cobra.Command{
	Use:   "align",
	Short: "Sets telescope alignment mode",
	Long: `Sets telescope to Land, Polar and AltAz alignment modes.
	
	Examples: scopectl align --mode land
	
	Usage: scopectl align --mode [mode]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("align called")
		log.Println("align called")
		mode, err := cmd.Flags().GetString("mode")
		if err != nil {
			fmt.Printf("error retrieving alignment mode: %s\n", err.Error())
			return err
		}
		if mode == "" {
			return errors.New("missing alignment mode")
		}
		log.Println("Mode Flag Value: " + mode)
		pmap := make(map[string]string)

		pmap["mode"] = mode

		scope_res, err := handlers.AlignCommandHandler(pmap)
		if err != nil {
			fmt.Printf("error calling server API server: %s\n", err.Error())
			return err
		}
		fmt.Printf("align command responded: %s\n", scope_res.Response)
		return nil
	},
}

func init() {
	alignCmd.Flags().StringP("mode", "m", "", "align mode")
	rootCmd.AddCommand(alignCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// alignCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// alignCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
