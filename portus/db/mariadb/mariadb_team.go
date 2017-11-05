package mariadb

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
	"strconv"
)

type TeamSearch struct {
	ID   int //TypeSearch=0
	Name string

	TypeSearch string
}

const TYPE_TEAM_ID = "id"
const TYPE_TEAM_NAME = "name"

func (m *Mariadb) GetTeam(s TeamSearch) (t Team, err error) {

	var strquery string
	strselect := "select id,name,hidden from teams where "
	switch s.TypeSearch {
	case TYPE_TEAM_ID:
		strquery = strselect + "id=" + strconv.Itoa(s.ID)
	case TYPE_TEAM_NAME:
		strquery = strselect + "name=" + s.Name
	}

	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return t, fmt.Errorf("ERROR:GetTeam -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&t.ID,
			&t.Name,
			&t.Hidden)

		if err != nil {
			return t, fmt.Errorf("ERROR:GetTeam -> Scan:%v", err)
		}
	}
	fmt.Println(t)
	return t, nil
}
