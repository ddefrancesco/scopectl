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
		Body: pmap["mode"],
	}
	var httpUrl string = viper.GetString("environments.test.url")
	var httpPort string = viper.GetString("environments.test.port")
	client := etxClient.NewClient(httpUrl+":"+httpPort, "POST", *etxRequestPath, *bodyRequest)

	scopeResponse, err := client.GetPost()
	if err != nil {
		return nil, err
	}

	return scopeResponse, nil
}
