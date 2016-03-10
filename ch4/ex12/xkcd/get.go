package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	Safe_title string
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func Get(number int) (*Comic, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://xkcd.com/"+strconv.Itoa(number)+"/info.0.json", nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("get failed: %s", resp.Status)
	}
	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
