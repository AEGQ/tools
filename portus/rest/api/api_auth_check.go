package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Res struct {
	Auth bool
}

// for portus_rmi tool
func (api Api) AuthCheck(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	var username string = params.Get(":username")
	var namespace string = params.Get(":namespace")

	ok, err := api.P.AuthCheck(username, namespace)
	if err != nil {
		fmt.Println("Error: AuthCheck -> AuthCheck(username, namespace)", err)
	}
	res := Res{Auth: ok}
	b, _ := json.Marshal(res)
	fmt.Fprintln(w, string(b))
}
