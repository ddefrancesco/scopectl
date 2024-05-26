/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package handlers

import (
	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/spf13/viper"
)

func AckCommandHandler() (*[]etxClient.ScopeResponse, error) {

	var etxRequestPath = &etxClient.RequestPath{
		Command: "ack",
		Items:   nil,
	}

	var bodyRequest = &etxClient.ScopeBodyRequest{
		Body: nil,
	}
	env := viper.GetString("environment")
	var httpUrl string = viper.GetString("environments." + env + ".url")
	var httpPort string = viper.GetString("environments." + env + ".port")
	client := etxClient.NewClient(httpUrl+":"+httpPort, "GET", etxRequestPath, bodyRequest)

	scopeResponse, err := client.Get()
	if err != nil {
		return nil, err
	}

	return scopeResponse, nil
}
