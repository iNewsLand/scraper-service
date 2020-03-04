package scraper

type projectsScrapyRequest struct {
	NodeName string `json:"node_name"`
	Status string `json:"status"`
	Projects []string `json:"projects"`
}
