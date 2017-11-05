package search

// ImageResult describes a search result returned from a registry
type ImageResult struct {
	// StarCount indicates the number of stars this repository has
	StarCount int `json:"star_count"`
	// IsOfficial is true if the result is from an official repository.
	IsOfficial bool `json:"is_official"`
	// Name is the name of the repository
	Name string `json:"name"`
	// IsAutomated indicates whether the result is automated
	IsAutomated bool `json:"is_automated"`
	// Description is a textual description of the repository
	Description string `json:"description"`
}

// ImageResults lists a collection search results returned from a registry
type ImageResults struct {
	// Query contains the query string that generated the search results
	Query string `json:"query"`
	// NumResults indicates the number of results the query returned
	NumResults int `json:"num_results"`
	// Results is a slice containing the actual results for the search
	Results []ImageResult `json:"results"`
}

// ImageResultsByStars sorts search results in descending order by number of stars.
type ImageResultsByStars []ImageResult

func (r ImageResultsByStars) Len() int           { return len(r) }
func (r ImageResultsByStars) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ImageResultsByStars) Less(i, j int) bool { return r[j].StarCount < r[i].StarCount }

// TagResult sss
type TagResult struct {
	Layer string `json:"layer,omitempty"`
	Name  string `json:"name"`
}
