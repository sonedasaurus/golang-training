package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// closeIssues queries the GitHub issue tracker.
func CloseIssues(number string, title string) error {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return err
	}
	var config Config
	json.Unmarshal(file, &config)
	token := base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(config.Username) + ":" + url.QueryEscape(config.Token)))

	client := &http.Client{}
	d := Data{
		title,
	}
	reqJSON, _ := json.Marshal(d)
	req, err := http.NewRequest("PATCH", EditIssuesURL+number, bytes.NewBuffer(reqJSON))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Basic "+token)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("edit Issues failed: %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}
