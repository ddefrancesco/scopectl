/*
Copyright © 2024 Daniele De Francesco ddefrancesco@gmail.com
*/
package handlers

import (
	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/spf13/viper"
)

func InfoCommandHandler(params map[string]string) (*[]etxClient.ScopeResponse, error) {
	//TODO InfoCommandHandler
	var etxRequestPath = &etxClient.RequestPath{
		Command: "info",
		Items:   params,
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
	for _, v := range *scopeResponse {
		//log.Printf("%s\n", v.Response)
		if v.Code != 200 {
			return nil, err
		}
	}

	return scopeResponse, nil
}
