package mariadb

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
)

func (m *Mariadb) GetRegistry() (r Registry, err error) {

	strquery := "select id,name,hostname,use_ssl from registries;"
	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return r, fmt.Errorf("ERROR:GetRegistry -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&r.ID,
			&r.Name,
			&r.Hostname,
			&r.Use_ssl)

		if err != nil {
			return r, fmt.Errorf("ERROR:GetRegistry -> Scan:%v", err)
		}

		//当前版本portus只支持一个registry
		break
	}
	return r, nil
}
