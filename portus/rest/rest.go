package rest

import (
	"fmt"
	"net/http"
	. "github.com/AEGQ/tools/portus/rest/api"
	"strconv"

	"github.com/drone/routes"
)

const (
	LISTENPORT   = 5050
	BASE_VERSION = "v1"
)

func StartRestServer() {
	api, err := NewApi()
	if err != nil {
		fmt.Println("ERROR:StartRestServer -> NewApi:%v", err)
	}

	fmt.Println("START...")

	mux := routes.New()

	mux.Get("/v1/repos", api.GetRepo)
	mux.Get("/v1/repos/username/:username", api.GetRepoFromUser)
	mux.Get("/v1/auth/:username/:namespace", api.AuthCheck)
	http.ListenAndServe(":"+strconv.Itoa(LISTENPORT), mux)
}
