package mariadb

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
	"strconv"
)

type NamespaceSearch struct {
	ID          int
	Name        string
	Team_id     int
	Registry_id int
	Global      int
	Visibility  int //0(namespace owner can pull),1(login users can pull),2(public)

	TypeSearch string
}

const TYPE_NAMESPACE_ID = "id"
const TYPE_NAMESPACE_NAME = "name"
const TYPE_NAMESPACE_TEAM_ID = "team_id"
const TYPE_NAMESPACE_REGISTRY_ID = "registry_id"
const TYPE_NAMESPACE_GLOBAL = "global"
const TYPE_NAMESPACE_VISIBILITY = "visibility"

func (m *Mariadb) GetNamespaces(s NamespaceSearch) (ns []NameSpace, err error) {

	var strquery string
	strselect := "select id,name,team_id,registry_id,global,visibility from namespaces where "

	switch s.TypeSearch {
	case TYPE_NAMESPACE_ID:
		strquery = strselect + "id=" + strconv.Itoa(s.ID)
	case TYPE_NAMESPACE_NAME:
		strquery = strselect + "name='" + s.Name + "'"
	case TYPE_NAMESPACE_TEAM_ID:
		strquery = strselect + "team_id=" + strconv.Itoa(s.Team_id)
	case TYPE_NAMESPACE_REGISTRY_ID:
		strquery = strselect + "registry_id=" + strconv.Itoa(s.Registry_id)

	case TYPE_NAMESPACE_GLOBAL:
		strquery = strselect + "global=" + strconv.Itoa(s.Global)

	case TYPE_NAMESPACE_VISIBILITY:
		strquery = strselect + "visibility=" + strconv.Itoa(s.Visibility)
	}
	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return ns, fmt.Errorf("ERROR:GetNamespaces -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		n := NameSpace{}
		err = rows.Scan(&n.ID,
			&n.Name,
			&n.Team_id,
			&n.Registry_id,
			&n.Global,
			&n.Visibility)

		if err != nil {
			return ns, fmt.Errorf("ERROR:GetNamespaces -> Scan:%v", err)
		}
		ns = append(ns, n)
	}
	return ns, nil
}
