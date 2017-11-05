package portus

import . "github.com/AEGQ/tools/portus/db/mariadb"

type Portus struct {
	DB *Mariadb
}

func NewPortus() (p *Portus, err error) {

	p = new(Portus)
	p.DB, err = NewMariadb("portus_development")
	if err != nil {
		return nil, err
	}
	return p, nil
}
