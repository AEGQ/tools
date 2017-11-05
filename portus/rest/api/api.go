package api

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/portus"
)

type Api struct {
	P *Portus
}

func NewApi() (api *Api, err error) {
	api = new(Api)
	api.P, err = NewPortus()
	if err != nil {
		return nil, fmt.Errorf("Error: NewApi -> NewPortus")
	}

	return api, nil
}
