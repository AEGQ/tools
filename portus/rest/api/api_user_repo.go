package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// for portus_rmi tool
func (api Api) GetRepoFromUser(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	var username string = params.Get(":username")

	res, err := api.P.GetRepoFromUser(username)
	if err != nil {
		fmt.Println("Error: AuthCheck -> GetRepoFromUser(username)", err)
	}
	b, _ := json.Marshal(res)
	fmt.Fprintln(w, string(b))
}
