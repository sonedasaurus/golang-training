package github

import "time"

const ListIssuesURL = "https://api.github.com/repos/sonedazaurus/golang-training/issues"
const MilestonesURL = "https://api.github.com/repos/sonedazaurus/golang-training/milestones"
const TeamURL = "https://api.github.com/repos/sonedazaurus/golang-training/teams"

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

type Milestone struct {
	Number      int
	HTMLURL     string `json:"html_url"`
	Title       string
	Description string
	State       string
	CreatedAt   time.Time `json:"created_at"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
