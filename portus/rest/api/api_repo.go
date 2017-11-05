package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api Api) GetRepo(w http.ResponseWriter, r *http.Request) {

	m, err := api.P.GetRepositories()
	if err != nil {
		fmt.Println("Error: GetRepo -> GetRepositories", err)
	}
	b, _ := json.Marshal(m)
	fmt.Fprintln(w, string(b))
}
