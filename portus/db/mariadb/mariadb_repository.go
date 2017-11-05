package mariadb

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/types"
	"strconv"
)

type RepoSearch struct {
	ID           int
	Namespace_id int

	TypeSearch string
}

const TYPE_REPO_ID = "id"
const TYPE_REPO_NAMESPACE_ID = "namespace_id"

func (m *Mariadb) GetRepositories(s RepoSearch) (rs []Repository, err error) {

	var strquery string
	strselect := `select id,name,namespace_id,marked from repositories where `
	switch s.TypeSearch {
	case TYPE_REPO_ID:
		strquery = strselect + `id=` + strconv.Itoa(s.ID)

	case TYPE_REPO_NAMESPACE_ID:
		strquery = strselect + `Namespace_id=` + strconv.Itoa(s.Namespace_id)
	}

	rows, err := m.dbconn.Query(strquery)
	if err != nil {
		return rs, fmt.Errorf("ERROR:GetRepositories -> Query:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		r := Repository{}
		err = rows.Scan(&r.ID,
			&r.Name,
			&r.Namespace_id,
			&r.Marked)

		if err != nil {
			return rs, fmt.Errorf("ERROR:GetRepositories -> Scan:%v", err)
		}
		rs = append(rs, r)
	}
	return rs, nil

}
