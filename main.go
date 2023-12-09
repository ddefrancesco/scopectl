/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package main

import (
	"fmt"

	"github.com/ddefrancesco/scopectl/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("scopeconfig")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading in config: ", err)
	}
	cmd.Execute()
}
