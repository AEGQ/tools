package portus

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/mariadb"
)

func (p Portus) AuthCheck(username, namespace string) (ok bool, err error) {

	if username == "" || namespace == "" {
		return false, fmt.Errorf("ERROR:AuthCheck -> Empty params")
	}

	//check admin
	us := UserSearch{Username: username, TypeSearch: TYPE_USER_USERNAME}
	u, err := p.DB.GetUsers(us)
	if err != nil {
		return false, fmt.Errorf("ERROR:AuthCheck -> GetUsers:%v", err)
	}
	if len(u) != 1 {
		return false, nil
	}
	if u[0].Admin == 1 {
		return true, nil
	}

	//check usernaem belongs to namespace'team
	ns := NamespaceSearch{Name: namespace, TypeSearch: TYPE_NAMESPACE_NAME}

	n, err := p.DB.GetNamespaces(ns)
	if err != nil {
		return false, fmt.Errorf("ERROR:AuthCheck -> GetNamespaces:%v", err)
	}
	if len(n) != 1 {
		return false, nil
	}
	tus := TeamUserSearch{Team_id: n[0].Team_id, TypeSearch: TYPE_TEAM_USER_TEAM_ID}
	tu, err := p.DB.GetTeamUsers(tus)
	if err != nil {
		return false, fmt.Errorf("ERROR:AuthCheck -> GetTeamUsers:%v", err)
	}

	for _, t := range tu {

		if t.User_id == u[0].ID {
			if t.Role == 2 { //owner
				return true, nil
			}
		}
	}

	return false, nil
}
