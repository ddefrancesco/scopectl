package handlers

import (
	etxClient "github.com/ddefrancesco/scopectl/restclient"
)

func AlignCommandHandler(pmap map[string]string) (*etxClient.ScopeResponse, error) {

	var etxRequestPath = &etxClient.RequestPath{
		Command: "align",
		Items:   pmap,
	}

	client := etxClient.NewClient("http://localhost:8000", "POST", *etxRequestPath)

	scopeResponse, err := client.GetPost()
	if err != nil {
		return nil, err
	}

	return scopeResponse, nil
}
