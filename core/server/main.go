package server

import (
	"mesh/core/ip"
	"mesh/core/mesh"
	"net/http"
)

type API struct {
	Route string
	Ip    *ip.Ip
	Code  string
}

func (api *API) Send(a mesh.Info) (error, *http.Response) {
	ipAddr := api.Ip.Address

	client := &http.Client{}
	req, _ := http.NewRequest("GET", ipAddr+"/"+api.Route, nil)
	req.Header.Set("MeshLAPICode", api.Code)
	resp, err := client.Do(req)

	if err != nil {
		return err, nil
	}
	return nil, resp
}