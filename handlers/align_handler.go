/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package handlers

import (
	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/spf13/viper"
)

func AlignCommandHandler(pmap map[string]string) (*etxClient.ScopeResponse, error) {

	var etxRequestPath = &etxClient.RequestPath{
		Command: "align",
		Items:   pmap,
	}

	var bodyRequest = &etxClient.ScopeBodyRequest{
		Body: pmap,
	}
	env := viper.GetString("environment")
	var httpUrl string = viper.GetString("environments." + env + ".url")
	var httpPort string = viper.GetString("environments." + env + ".port")
	client := etxClient.NewClient(httpUrl+":"+httpPort, "POST", etxRequestPath, bodyRequest)

	scopeResponse, err := client.Post()
	if err != nil {
		return nil, err
	}
	if scopeResponse.Code != 202 {
		return nil, err
	}

	return scopeResponse, nil
}
