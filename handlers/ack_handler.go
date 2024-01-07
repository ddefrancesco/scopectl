package handlers

import (
	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/spf13/viper"
)

func AckCommandHandler() (*etxClient.ScopeResponse, error) {

	var etxRequestPath = &etxClient.RequestPath{
		Command: "ack",
		Items:   nil,
	}

	var bodyRequest = &etxClient.ScopeBodyRequest{
		Body: "",
	}
	var httpUrl string = viper.GetString("environments.test.url")
	var httpPort string = viper.GetString("environments.test.port")
	client := etxClient.NewClient(httpUrl+":"+httpPort, "GET", *etxRequestPath, *bodyRequest)

	scopeResponse, err := client.GetPost()
	if err != nil {
		return nil, err
	}

	return scopeResponse, nil
}
