package cmd

import (
	"log"

	"github.com/ddefrancesco/scopectl/handlers"
	"github.com/spf13/cobra"
)

// gotoCmd represents the settings command
var gotoCmd = &cobra.Command{
	Use:   "goto",
	Short: "slew telescope to target",
	Long: `Slew telescope to target.	
	
	Examples: scopectl goto --ngc [id_ngc]
	
	Parameters:

	-a  --altitude [altitude] 
	--target_ra [RA] 
	--target_dec [DEC] 
	--ngc [NGC]
    --target_azimuth [azimuth]

	Usage: scopectl goto --ngc=NGC0221
		  `,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println("goto command called")

		pmap := make(map[string]string)

		pmap["altitude"], _ = cmd.Flags().GetString("altitude")
		pmap["current_date"], _ = cmd.Flags().GetString("current_date")
		pmap["target_dec"], _ = cmd.Flags().GetString("target_dec")
		pmap["local_time"], _ = cmd.Flags().GetString("local_time")
		pmap["target_ra"], _ = cmd.Flags().GetString("target_ra")
		pmap["local_sidereal_time"], _ = cmd.Flags().GetString("local_sidereal_time")
		pmap["ngc"], _ = cmd.Flags().GetString("ngc")
		pmap["hip"], _ = cmd.Flags().GetString("hip")
		pmap["goto"], _ = cmd.Flags().GetString("goto")
	scope_res, err := gotoHandler(pmap)
		if err != nil {
			log.Printf("error calling server API server: %s\n", err.Error())
			return err
		}
		log.Printf("goto command responded: %s\n", scope_res.Response)
		return nil
	},
}

// gotoHandler is a package-level variable so tests can replace it.
var gotoHandler = handlers.GotoCommandHandler

func init() {

	gotoCmd.Flags().StringP("altitude", "a", "", "Target object altitude to sDD*MM# or sDD*MM’SS# [LX 16”, Autostar, LX200GPS/RCX400]")
	gotoCmd.Flags().StringP("current_date", "C", "", "Change Handbox Date to MM/DD/YY")
	gotoCmd.Flags().StringP("target_dec", "d", "", "Set target object declination to sDD*MM or sDD*MM:SS depending on the current precision setting")
	gotoCmd.Flags().StringP("local_time", "L", "", "Set the local Time")
	gotoCmd.Flags().StringP("target_ra", "r", "", "Set target object RA to HH:MM.T or HH:MM:SS depending on the current precision setting")
	gotoCmd.Flags().StringP("local_sidereal_time", "S", "", "Sets the local sidereal time to HH:MM:SS")
	gotoCmd.Flags().StringP("ngc", "M", "", "Slew to the target object specified by its NGC number")
	gotoCmd.Flags().StringP("hip", "H", "", "Slew to the target object specified by its Hipparcos number")
	gotoCmd.Flags().StringP("solar_sys", "s", "", "Slew to the target object specified by its Solar System body name")
	gotoCmd.Flags().StringP("goto", "g", "", "Temporary alias for --ngc")
	rootCmd.AddCommand(gotoCmd)

}
