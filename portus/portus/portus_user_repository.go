package portus

import (
	"fmt"
	. "github.com/AEGQ/tools/portus/db/mariadb"
)

func (p Portus) GetRepoFromUser(username string) (repos map[string][]Repo, err error) {

	repos = make(map[string][]Repo)
	if username == "" {
		return repos, fmt.Errorf("ERROR:GetRepoFromUser -> Empty params")
	}

	//get user_id
	us := UserSearch{Username: username, TypeSearch: TYPE_USER_USERNAME}
	u, err := p.DB.GetUsers(us)
	if err != nil {
		return repos, fmt.Errorf("ERROR:GetRepoFromUser -> GetUsers:%v", err)
	}
	if len(u) != 1 {
		return repos, fmt.Errorf("ERROR:GetRepoFromUser -> No such username.")
	}

	//get team_id
	tus := TeamUserSearch{User_id: u[0].ID, TypeSearch: TYPE_TEAM_USER_USER_ID}
	tu, err := p.DB.GetTeamUsers(tus)
	if err != nil {
		return repos, fmt.Errorf("ERROR:GetRepoFromUser -> GetTeamUsers:%v", err)
	}

	//get namespaces
	for _, v := range tu {
		ns := NamespaceSearch{Team_id: v.Team_id, TypeSearch: TYPE_NAMESPACE_TEAM_ID}
		n, err := p.DB.GetNamespaces(ns)
		if err != nil {
			return repos, fmt.Errorf("ERROR:GetRepoFromUser -> GetNamespaces:%v", err)
		}

		//get repos
		for _, na := range n {
			if na.Visibility != 2 { // only public namespace
				continue
			}

			s := RepoSearch{Namespace_id: na.ID, TypeSearch: TYPE_REPO_NAMESPACE_ID}
			rs, err := p.DB.GetRepositories(s)
			if err != nil {
				return repos, fmt.Errorf("Error:GetRepoFromUser -> GetRepositories:", err)
			}
			//get tags
			var repo_arry []Repo
			for _, r := range rs {
				var rp Repo
				//all tags with each repo
				ts, err := p.DB.GetTags(r.ID)
				if err != nil {
					return repos, fmt.Errorf("Error:GetRepoFromUser -> GetTags:", err)
				}
				//return only name and tags
				rp.Name = r.Name
				for _, t := range ts {
					rp.Tags = append(rp.Tags, t.Name)
				}
				repo_arry = append(repo_arry, rp)
			}
			repos[na.Name] = repo_arry
		}
	}
	return repos, nil
}
