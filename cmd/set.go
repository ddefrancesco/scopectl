/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package cmd

import (
	"log"

	"github.com/ddefrancesco/scopectl/handlers"
	"github.com/spf13/cobra"
)

// setCmd represents the settings command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets telescope settings",
	Long: `Sets telescope settings, like target altitude, azimuth, declination, etc.
	
	Examples: scopectl set --altitude [altitude] 
	
	Usage: scopectl set --<setting1> [value1] --<setting2> [value2] ...
		  `,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println("set command called")

		pmap := make(map[string]string)

		pmap["altitude"], _ = cmd.Flags().GetString("altitude")
		pmap["brighter_limit"], _ = cmd.Flags().GetString("brighter_limit")
		pmap["baud_speed"], _ = cmd.Flags().GetString("baud_speed")
		pmap["current_date"], _ = cmd.Flags().GetString("current_date")
		pmap["target_dec"], _ = cmd.Flags().GetString("target_dec")
		pmap["selenographic_lat"], _ = cmd.Flags().GetString("selenographic_lat")
		pmap["selenographic_long"], _ = cmd.Flags().GetString("selenographic_long")
		pmap["fainter_limit"], _ = cmd.Flags().GetString("fainter_limit")
		pmap["field_diameter"], _ = cmd.Flags().GetString("field_diameter")
		pmap["current_site_long"], _ = cmd.Flags().GetString("current_site_long")
		pmap["utc_offset"], _ = cmd.Flags().GetString("utc_offset")
		pmap["dst"], _ = cmd.Flags().GetString("dst")
		pmap["min_elev_limit"], _ = cmd.Flags().GetString("min_elev_limit")
		pmap["smallest_limit"], _ = cmd.Flags().GetString("smallest_limit")
		pmap["local_time"], _ = cmd.Flags().GetString("local_time")
		pmap["high_limit"], _ = cmd.Flags().GetString("high_limit")
		pmap["quality"], _ = cmd.Flags().GetString("quality")
		pmap["target_ra"], _ = cmd.Flags().GetString("target_ra")
		pmap["largest_limit"], _ = cmd.Flags().GetString("largest_limit")
		pmap["local_sidereal_time"], _ = cmd.Flags().GetString("local_sidereal_time")
		pmap["current_site_lat"], _ = cmd.Flags().GetString("current_site_lat")
		pmap["tracking_rate"], _ = cmd.Flags().GetString("tracking_rate")
		pmap["max_slew_rate"], _ = cmd.Flags().GetString("max_slew_rate")
		pmap["target_azimuth"], _ = cmd.Flags().GetString("target_azimuth")

		scope_res, err := handlers.SetCommandHandler(pmap)
		if err != nil {
			log.Printf("error calling server API server: %s\n", err.Error())
			return err
		}
		log.Printf("set command responded: %s\n", scope_res.Response)
		return nil
	},
}

func init() {

	setCmd.Flags().StringP("altitude", "a", "", "Target object altitude to sDD*MM# or sDD*MM’SS# [LX 16”, Autostar, LX200GPS/RCX400]")
	setCmd.Flags().StringP("brighter_limit", "b", "", "Brighter limit to the ASCII decimal magnitude string")
	setCmd.Flags().StringP("baud_speed", "B", "", "Set Baud Rate n, where n is an ASCII digit (1..9")
	setCmd.Flags().StringP("current_date", "C", "", "Change Handbox Date to MM/DD/YY")
	setCmd.Flags().StringP("target_dec", "d", "", "Set target object declination to sDD*MM or sDD*MM:SS depending on the current precision setting")
	setCmd.Flags().StringP("selenographic_lat", "E", "", "Sets target object to the specificed selenographic latitude on the Moon.")
	setCmd.Flags().StringP("selenographic_long", "e", "", "Sets the target object to the specified selenogrphic longitude on the Moon")
	setCmd.Flags().StringP("fainter_limit", "f", "", "Set faint magnitude limit to sMM.M")
	setCmd.Flags().StringP("field_diameter", "F", "", "Set FIELD/IDENTIFY field diameter to NNN arc minutes")
	setCmd.Flags().StringP("current_site_long", "g", "", "Set current site’s longitude to DDD*MM an ASCII position string")
	setCmd.Flags().StringP("utc_offset", "G", "", "Set the number of hours added to local time to yield UTC")
	setCmd.Flags().StringP("dst", "H", "", "Set Daylight Savings Mode [Autostar II Only]. D=1 Sets Daylight savings. D=0 Clears Daylight savings")
	setCmd.Flags().StringP("min_elev_limit", "m", "", "Set the minimum object elevation limit to DD#")
	setCmd.Flags().StringP("smallest_limit", "l", "", "Set the size of the smallest object returned by FIND/BROWSE to NNNN arc minutes")
	setCmd.Flags().StringP("local_time", "L", "", "Set the local Time")
	setCmd.Flags().StringP("high_limit", "o", "", "Set highest elevation to which the telescope will slew")
	setCmd.Flags().StringP("quality", "q", "", "Step the quality of limit used in FIND/BROWSE through its cycle of VP … SU. Current setting can be queried with :Gq#")
	setCmd.Flags().StringP("target_ra", "r", "", "Set target object RA to HH:MM.T or HH:MM:SS depending on the current precision setting")
	setCmd.Flags().StringP("largest_limit", "s", "", "Set the size of the largest object the FIND/BROWSE command will return to NNNN arc minutes")
	setCmd.Flags().StringP("local_sidereal_time", "S", "", "Sets the local sidereal time to HH:MM:SS")
	setCmd.Flags().StringP("current_site_lat", "t", "", "Sets the current site latitude to sDD*MM#")
	setCmd.Flags().StringP("tracking_rate", "T", "", "Sets the current tracking rate to TTT.T hertz, assuming a model where a 60.0 Hertz synchronous motor will cause the RA axis to make exactly one revolution in 24 hours.")
	setCmd.Flags().StringP("max_slew_rate", "w", "", "Set maximum slew rate to N degrees per second. N is the range (2..8)")
	setCmd.Flags().StringP("target_azimuth", "z", "", "Sets the target Object Azimuth [LX 16” and LX200GPS/RCX400 only]")

	rootCmd.AddCommand(setCmd)

}
