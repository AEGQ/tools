package mariadb

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
	"strconv"
)

type TeamUserSearch struct {
	ID      int //TypeSearch=0
	User_id int
	Team_id int

	TypeSearch string
}

const TYPE_TEAM_USER_ID = "id"
const TYPE_TEAM_USER_USER_ID = "user_id"
const TYPE_TEAM_USER_TEAM_ID = "team_id"

func (m *Mariadb) GetTeamUsers(s TeamUserSearch) (ts []TeamUser, err error) {

	var strquery string
	strselect := "select id,user_id,team_id,role from team_users where "
	switch s.TypeSearch {
	case TYPE_TEAM_USER_ID:
		strquery = strselect + "id=" + strconv.Itoa(s.ID)
	case TYPE_TEAM_USER_USER_ID:
		strquery = strselect + "user_id=" + strconv.Itoa(s.User_id)
	case TYPE_TEAM_USER_TEAM_ID:
		strquery = strselect + "team_id=" + strconv.Itoa(s.Team_id)
	}

	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return ts, fmt.Errorf("ERROR:GetTeamUsers -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		t := TeamUser{}
		err = rows.Scan(&t.ID,
			&t.User_id,
			&t.Team_id,
			&t.Role)

		if err != nil {
			return ts, fmt.Errorf("ERROR:GetTeamUsers -> Scan:%v", err)
		}
		ts = append(ts, t)
	}
	return ts, nil
}
