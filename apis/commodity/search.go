package commodity

type SearchQuery struct {
	Search   string `json:"search"`
	Accurate bool   `json:"accurate"`
	Range    string `json:"range"`
}
