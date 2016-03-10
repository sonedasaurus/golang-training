package github

import "time"

const SearchIssuesURL = "https://api.github.com/search/issues"
const CreateIssuesURL = "https://api.github.com/repos/sonedazaurus/golang-training/issues"
const ListIssuesURL = "https://api.github.com/repos/sonedazaurus/golang-training/issues"
const EditIssuesURL = "https://api.github.com/repos/sonedazaurus/golang-training/issues/"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Config struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Data struct {
	Title string `json:"title"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
