package item

type SearchQuery struct {
	Search   string `json:"search"`
	Accurate bool   `json:"accurate"`
	Range    string `json:"range"`
}
