package server

import (
	"io"
	"mesh/core/ip"
	"net/http"
)

type API struct {
	Route string
	Ip    *ip.Ip
	Code  string
}

func (api *API) Send(body io.ReadCloser, header http.Header) (error, *http.Response) {
	ipAddr := api.Ip.Address

	client := &http.Client{}
	req, _ := http.NewRequest("GET", ipAddr+"/"+api.Route, body)
	req.Header.Set("MeshLAPICode", api.Code)
	for key, values := range header {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		return err, nil
	}
	return nil, resp
}
