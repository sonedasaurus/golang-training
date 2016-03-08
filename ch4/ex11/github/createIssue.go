package github

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Data struct {
	Title string `json:"title"`
}

// CreateIssues queries the GitHub issue tracker.
func CreateIssues(title string) error {
	file, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return err
	}
	var config Config
	json.Unmarshal(file, &config)
	token := base64.StdEncoding.EncodeToString([]byte(url.QueryEscape(config.Username) + ":" + url.QueryEscape(config.Token)))

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	d := Data{
		title,
	}
	reqJSON, _ := json.Marshal(d)
	req, err := http.NewRequest("POST", CreateIssuesURL, bytes.NewBuffer(reqJSON))
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
		return fmt.Errorf("create Issues failed: %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}
