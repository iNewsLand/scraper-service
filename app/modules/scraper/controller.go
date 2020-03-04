package scraper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var url = os.Getenv("SCRAPYD_URL")

func getProjects() ([]string, error) {
	prefix := "listprojects.json"
	req, err := http.Get(url + prefix)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var projects projectsScrapyRequest
	err = json.Unmarshal(body, projects)
	if err != nil {
		return nil, err
	}
	var projectsName []string
	for _, projectName := range projects.Projects {
		projectsName = append(projectsName, projectName)
	}
	return projectsName, nil
}
