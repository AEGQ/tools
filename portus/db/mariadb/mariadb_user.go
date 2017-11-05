package mariadb

import (
	"database/sql"
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
	"strconv"
)

type UserSearch struct {
	ID           int //TypeSearch=0
	Username     string
	Admin        int
	Namespace_id int //TypeSearch=3

	TypeSearch string
}

const TYPE_USER_ID = "id"
const TYPE_USER_USERNAME = "username"
const TYPE_USER_ADMIN = "admin"
const TYPE_USER_NAMESPACE_ID = "namespace_id"

func (m *Mariadb) GetUsers(s UserSearch) (us []User, err error) {

	var strquery string
	strselect := "select id,username,admin,enabled,ldap_name,namespace_id from users where "
	switch s.TypeSearch {
	case TYPE_USER_ID:
		strquery = strselect + "id=" + strconv.Itoa(s.ID)
	case TYPE_USER_USERNAME:
		strquery = strselect + "username='" + s.Username + "'"
	case TYPE_USER_ADMIN:
		strquery = strselect + "admim=" + strconv.Itoa(s.Admin)
	case TYPE_USER_NAMESPACE_ID:
		strquery = strselect + "namespace_id=" + strconv.Itoa(s.Namespace_id)

	}

	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return us, fmt.Errorf("ERROR:GetUsers -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		u := User{}
		var i sql.NullInt64
		var l sql.NullString
		err = rows.Scan(&u.ID,
			&u.Username,
			&u.Admin,
			&u.Enabled,
			&l,
			&i)
		if l.Valid {
			u.Ldap_name = l.String
		} else {
			u.Ldap_name = ""
		}
		if i.Valid {
			u.Namespace_id = int(i.Int64)
		} else {
			u.Namespace_id = 0
		}

		if err != nil {
			return us, fmt.Errorf("ERROR:GetUsers -> Scan:%v", err)
		}
		us = append(us, u)
	}
	return us, nil
}
