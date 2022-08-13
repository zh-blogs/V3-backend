package types

// BlogSimple ...
type BlogSimple struct {
	ID   string   `json:"id"`
	Idx  int      `json:"idx"`
	Name string   `json:"name"`
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
	Sign string   `json:"sign"`
	Feed string   `json:"feed"`
}
