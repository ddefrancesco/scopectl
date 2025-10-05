/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/ddefrancesco/scopectl/handlers"
	"github.com/spf13/cobra"
)

type Info string
type InfoCommandValue string

const (
	InfoAltitude          Info = "current_altitude"
	InfoLTT               Info = "local_telescope_time"
	InfoBrighterMagLimit  Info = "browse_bml"
	InfoCurrentDate       Info = "current_date"
	InfoClockFmt          Info = "clock_fmt"
	InfoDeclination       Info = "scope_declination"
	InfoSelectedTargetDec Info = "target_dec"
	InfoFieldDiameter     Info = "scope_field_diameter"
	InfoFainterMagLimit   Info = "fainter_mag_limit"
	InfoUTCOffset         Info = "utc_offset"
	InfoCurrentSiteLong   Info = "site_long"
	InfoHighLimit         Info = "slew_high_limit"
	InfoLocalTime24h      Info = "current_local_time_24h"
	InfoLargerSizeLimit   Info = "view_larger_size_limit"
	InfoLowerSizeLimit    Info = "view_lower_size_limit"
	InfoMinimumQuality    Info = "minimum_find_quality"
	InfoRA                Info = "scope_ra"
	InfoCurrentTargetRA   Info = "target_ra"
	InfoSiderealTime      Info = "sidereal_time"
	InfoSmallerSizeLimit  Info = "view_smaller_size_limit"
	InfoTrackingRate      Info = "scope_tracking_rate"
	InfoCurrentSiteLat    Info = "site_lat"
	InfoFirmwareDate      Info = "firmware_date"
	InfoFirmwareVersion   Info = "firmware_version"
	InfoProductName       Info = "product_name"
	InfoFirmwareTime      Info = "firmware_time"
	InfoDeepsky           Info = "deepsky"
	InfoAzimuth           Info = "current_azimuth"
)

const (
	InfoAltitudeCmd          InfoCommandValue = ":GA#"
	InfoLTTCmd               InfoCommandValue = ":Ga#"
	InfoBrighterMagLimitCmd  InfoCommandValue = ":Gb#"
	InfoCurrentDateCmd       InfoCommandValue = ":GC#"
	InfoClockFmtCmd          InfoCommandValue = ":Gc#"
	InfoDeclinationCmd       InfoCommandValue = ":GD#"
	InfoSelectedTargetDecCmd InfoCommandValue = ":Gd#"
	InfoFieldDiameterCmd     InfoCommandValue = ":GF#"
	InfoFainterMagLimitCmd   InfoCommandValue = ":Gf#"
	InfoUTCOffsetCmd         InfoCommandValue = ":GG#"
	InfoCurrentSiteLongCmd   InfoCommandValue = ":Gg#"
	InfoHighLimitCmd         InfoCommandValue = ":Gh#"
	InfoLocalTime24hCmd      InfoCommandValue = ":GL#"
	InfoLargerSizeLimitCmd   InfoCommandValue = ":GI#"
	InfoLowerSizeLimitCmd    InfoCommandValue = ":Go#"
	InfoMinimumQualityCmd    InfoCommandValue = ":Gq#"
	InfoRACmd                InfoCommandValue = ":GR#"
	InfoCurrentTargetRACmd   InfoCommandValue = ":Gr#"
	InfoSiderealTimeCmd      InfoCommandValue = ":GS#"
	InfoSmallerSizeLimitCmd  InfoCommandValue = ":Gs#"
	InfoTrackingRateCmd      InfoCommandValue = ":GT#"
	InfoCurrentSiteLatCmd    InfoCommandValue = ":Gt#"
	InfoFirmwareDateCmd      InfoCommandValue = ":GVD#"
	InfoFirmwareVersionCmd   InfoCommandValue = ":GVN#"
	InfoProductNameCmd       InfoCommandValue = ":GVP#"
	InfoFirmwareTimeCmd      InfoCommandValue = ":GVT#"
	InfoDeepskyCmd           InfoCommandValue = ":Gy#"
	InfoAzimuthCmd           InfoCommandValue = ":GZ#"
)

func initCommandDictionary() map[InfoCommandValue]Info {
	infoMap := make(map[InfoCommandValue]Info)
	infoMap[InfoAltitudeCmd] = InfoAltitude
	infoMap[InfoLTTCmd] = InfoLTT
	infoMap[InfoBrighterMagLimitCmd] = InfoBrighterMagLimit
	infoMap[InfoCurrentDateCmd] = InfoCurrentDate
	infoMap[InfoClockFmtCmd] = InfoClockFmt
	infoMap[InfoDeclinationCmd] = InfoDeclination
	infoMap[InfoSelectedTargetDecCmd] = InfoSelectedTargetDec
	infoMap[InfoFieldDiameterCmd] = InfoFieldDiameter
	infoMap[InfoFainterMagLimitCmd] = InfoFainterMagLimit
	infoMap[InfoUTCOffsetCmd] = InfoUTCOffset
	infoMap[InfoCurrentSiteLongCmd] = InfoCurrentSiteLong
	infoMap[InfoHighLimitCmd] = InfoHighLimit
	infoMap[InfoLocalTime24hCmd] = InfoLocalTime24h
	infoMap[InfoLargerSizeLimitCmd] = InfoLargerSizeLimit
	infoMap[InfoLowerSizeLimitCmd] = InfoLowerSizeLimit
	infoMap[InfoMinimumQualityCmd] = InfoMinimumQuality
	infoMap[InfoRACmd] = InfoRA
	infoMap[InfoCurrentTargetRACmd] = InfoCurrentTargetRA
	infoMap[InfoSiderealTimeCmd] = InfoSiderealTime
	infoMap[InfoSmallerSizeLimitCmd] = InfoSmallerSizeLimit
	infoMap[InfoTrackingRateCmd] = InfoTrackingRate
	infoMap[InfoCurrentSiteLatCmd] = InfoCurrentSiteLat
	infoMap[InfoFirmwareDateCmd] = InfoFirmwareDate
	infoMap[InfoFirmwareVersionCmd] = InfoFirmwareVersion
	infoMap[InfoProductNameCmd] = InfoProductName
	infoMap[InfoFirmwareTimeCmd] = InfoFirmwareTime
	infoMap[InfoDeepskyCmd] = InfoDeepsky
	infoMap[InfoAzimuthCmd] = InfoAzimuth
	return infoMap
}

const infoTemplateWithSubsections = `
==============================================================================
Telescope Information
==============================================================================

{{- if . }}

Current Status:
--------------
{{- range . }}
{{- if or (hasPrefix .Cmd "current") (eq .Cmd "altitude") (eq .Cmd "azimuth") (eq .Cmd "declination") (eq .Cmd "ra") }}
  {{.Cmd}}: {{.Response}}
{{- end }}
{{- end }}

Target Information:
------------------
{{- range . }}
{{- if hasPrefix .Cmd "target" }}
  {{.Cmd}}: {{.Response}}
{{- end }}
{{- end }}

Site Information:
----------------
{{- range . }}
{{- if or (hasPrefix .Cmd "site") (eq .Cmd "utc_offset") (eq .Cmd "sidereal_time") }}
  {{.Cmd}}: {{.Response}}
{{- end }}
{{- end }}

Telescope Settings:
------------------
{{- range . }}
{{- if and (not (hasPrefix .Cmd "current")) (not (hasPrefix .Cmd "firmware")) (not (hasPrefix .Cmd "target")) (not (hasPrefix .Cmd "site")) (ne .Cmd "utc_offset") (ne .Cmd "sidereal_time") }}
  {{.Cmd}}: {{.Response}}
{{- end }}
{{- end }}

Firmware Information:
--------------------
{{- range . }}
{{- if  or (hasPrefix .Cmd "product") (hasPrefix .Cmd "firmware")}}
  {{.Cmd}}: {{.Response}}
{{- end }}
{{- end }}

{{- else }}
No telescope information available.
{{- end }}

==============================================================================
`

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

		// Usage
		//t := template.Must(template.New("info").Parse(infoTemplate))
		funcMap := template.FuncMap{
			"hasPrefix": strings.HasPrefix,
		}
		t := template.Must(template.New("info").Funcs(funcMap).Parse(infoTemplateWithSubsections))

		var cmdDescMap = initCommandDictionary()
		for i, v := range *scope_res {
			//log.Printf("info: %s\n", v.Response)
			v.Cmd = string(cmdDescMap[InfoCommandValue(v.Cmd)])
			(*scope_res)[i] = v
		}

		//Sort by Command
		sort.Slice(*scope_res, func(i, j int) bool {
			return (*scope_res)[i].Cmd < (*scope_res)[j].Cmd
		})
		err = t.Execute(os.Stdout, scope_res)
		if err != nil {
			log.Printf("error executing template: %v", err)
			return nil
		}
		return nil
	}}

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
