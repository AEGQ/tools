package portus

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/mariadb"
)

type Repo struct {
	Name string
	Tags []string
}

func (p Portus) GetRepositories() (repos map[string][]Repo, err error) {

	repos = make(map[string][]Repo)
	//all public namespace
	sp := NamespaceSearch{Visibility: 2, TypeSearch: TYPE_NAMESPACE_VISIBILITY}
	ns, err := p.DB.GetNamespaces(sp)
	if err != nil {
		return repos, fmt.Errorf("Error: GetRepositories -> GetNamespaces:", err)
	}
	for _, n := range ns {
		//all repos with each namespace
		s := RepoSearch{Namespace_id: n.ID, TypeSearch: TYPE_REPO_NAMESPACE_ID}
		rs, err := p.DB.GetRepositories(s)
		if err != nil {
			return repos, fmt.Errorf("Error: GetRepositories -> GetRepositories:", err)
		}
		var repo_arry []Repo
		for _, r := range rs {
			var rp Repo
			//all tags with each repo
			ts, err := p.DB.GetTags(r.ID)
			if err != nil {
				return repos, fmt.Errorf("Error: GetRepositories -> GetTags:", err)
			}
			//return only name and tags
			rp.Name = r.Name
			for _, t := range ts {
				rp.Tags = append(rp.Tags, t.Name)
			}
			repo_arry = append(repo_arry, rp)
		}
		repos[n.Name] = repo_arry
	}
	return repos, nil
}
