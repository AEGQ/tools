package portus

import (
	. "github.com/AEGQ/tools/portus/db/types"
)

func (p Portus) GetRegistry() (Registry, error) {

	//registry
	return p.DB.GetRegistry()
}
