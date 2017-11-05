package mariadb

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
	"strconv"
)

func (m *Mariadb) GetTags(repository_id int) (ts []Tag, err error) {

	strquery := "select id,name,repository_id,user_id,marked from tags where repository_id=" + strconv.Itoa(repository_id)

	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return ts, fmt.Errorf("ERROR:GetTags -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		t := Tag{}
		err = rows.Scan(&t.ID,
			&t.Name,
			&t.Repository_id,
			&t.User_id,
			&t.Marked)

		if err != nil {
			return ts, fmt.Errorf("ERROR:GetTags -> Scan:%v", err)
		}
		ts = append(ts, t)
	}
	return ts, nil
}
