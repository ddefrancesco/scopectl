/*
Copyright © 2023 Daniele De Francesco ddefrancesco@gmail.com
*/
package handlers

import (
	"github.com/ddefrancesco/scopectl/handlers/commons"
	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/spf13/viper"
)

// GotoCommandHandler handles the "goto" command logic.
// It builds the request body and sends it to the server using the REST client.
func GotoCommandHandler(pmap map[string]string) (*etxClient.ScopeResponse, error) {
	bodyMap := commons.ToScopeBodyMap(pmap)
	var etxRequestPath = &etxClient.RequestPath{
		Command: "move",
		Items:   bodyMap,
	}

	var bodyRequest = &etxClient.ScopeBodyRequest{
		Body: bodyMap,
	}
	env := viper.GetString("environment")
	httpUrl := viper.GetString("environments." + env + ".url")
	httpPort := viper.GetString("environments." + env + ".port")
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
