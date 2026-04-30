package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Respo struct {
	Login               string `json:"login"`
	Id                  int
	Node_id             string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Followers_url       string
	Following_url       string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Resp_type           string `json:"type"`
	User_view_type      string
	Site_admin          bool
	Name                string
	Company             string
	Blog                string
	Location            string
	Email               string
	Hireable            bool
	Bio                 string
	Twitter_username    string
	Public_repos        int
	Public_gists        int
	Followers           int
	Following           int
	Created_at          string
	Updated_at          string
}

/*
Hit any public API like https://api.github.com/users/torvalds
Decode the response into a struct
Library: net/http, encoding/json
Refreshes: http.Get, json.NewDecoder(resp.Body).Decode(), defer resp.Body.Close().
*/
func main() {
	resp, err := http.Get("https://api.github.com/users/torvalds")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var respo Respo

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("bad status: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&respo)
	if err != nil {
		log.Fatal(err)
	}

}
