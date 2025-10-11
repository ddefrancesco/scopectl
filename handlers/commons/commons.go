package commons

import (
	etxClient "github.com/ddefrancesco/scopectl/restclient"
)

func ToScopeBodyMap(pmap map[string]string) etxClient.BodyMap {
	_map := make(map[string]string)
	for k, v := range pmap {

		if v != "" {
			_map[k] = v
		}
	}
	return _map

}
