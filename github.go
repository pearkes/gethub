package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// type Repo represents a single repository
type Repo struct {
	FullName string `json:"full_name"`
	SSHUrl   string `json:"ssh_url"`
	HTTPSUrl string `json:"clone_url"`
}

func (r Repo) Owner() string {
	return strings.Split(r.FullName, "/")[0]
}

func (r Repo) Name() string {
	return strings.Split(r.FullName, "/")[1]
}

// Retrieves a list of all available repostories and builds them up into
// something we can handle locally. After this occurs, we begin our
// clone / fetch sequence.
func listRemoteRepostories(env Env) []Repo {
	log.Println("Retrieving remote repositories...")
	client := &http.Client{}

	req, err := http.NewRequest("GET",
		"https://api.github.com/user/repos?type=all&per_page=100", nil)

	req.SetBasicAuth(env.Config.Token, "")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var repos []Repo

	err = json.Unmarshal(body, &repos)

	if err != nil {
		fmt.Println(err)
	}

	log.Println(resp.Status)

	if resp.StatusCode != 200 {
		fmt.Println(red + "Uh oh, there was an error getting your repositories from GitHub. Here's what we got back:\n" + clear)
		fmt.Println(string(body))
		os.Exit(1)
	}

	log.Println(len(repos), "repositories retrieved from GitHub")

	return repos
}
