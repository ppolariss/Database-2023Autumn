package platform

type CreatePlatformRequest struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Country string `json:"country"`
}
