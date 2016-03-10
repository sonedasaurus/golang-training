package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ListIssues queries the GitHub issue tracker.
func ListIssues() ([]Issue, error) {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return nil, err
	}
	var config Config
	json.Unmarshal(file, &config)

	client := &http.Client{}
	req, err := http.NewRequest("GET", ListIssuesURL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(config.Username, config.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("list Issues failed: %s", resp.Status)
	}
	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
