package steps

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/multistep"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

// type Org represents an organization
type Org struct {
	Name string `json:"login"`
}

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

type StepRetrieveRepositories struct{}

func (*StepRetrieveRepositories) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Retrieving remote repositories...")
	var allRepos []Repo
	var endpoints []string

	token := state["token"].(string)

	fmt.Printf("Contacting GitHub... ")

	// Retrieve Organizations
	body := apiRequest(token, "/user/orgs")
	var orgs []Org
	err := json.Unmarshal(body, &orgs)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(len(orgs), "organizations retrieved from GitHub")

	// Build the endpoint for each organization
	for _, org := range orgs {
		endpoints = append(endpoints, "/orgs/"+org.Name+"/repos?type=all&per_page=100")
	}

	// Add the user's repos to the endpoint
	endpoints = append(endpoints, "/user/repos?type=all&per_page=100")

	var wg sync.WaitGroup
	// Asynchronously retrieve all repositories from GitHub
	for _, endpoint := range endpoints {
		wg.Add(1)

		go func(endpoint string) {
			repos := []Repo{}
			body := apiRequest(token, endpoint)
			err := json.Unmarshal(body, &repos)
			if err != nil {
				fmt.Println(err)
			}
			// Add the requested repos to the list of all repos
			allRepos = append(allRepos, repos...)
			// This one is done!
			wg.Done()
		}(endpoint)
	}

	// Wait for every endpoint to be requested
	wg.Wait()

	log.Println(len(allRepos), "repositories retrieved from GitHub")
	state["repos"] = allRepos

	fmt.Printf("%sdone%s\n", GREEN, CLEAR)

	return multistep.ActionContinue
}

func (*StepRetrieveRepositories) Cleanup(map[string]interface{}) {}

func apiRequest(token string, endpoint string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET",
		"https://api.github.com"+endpoint, nil)

	req.SetBasicAuth(token, "")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	log.Println(resp.Status)

	if resp.StatusCode != 200 {
		fmt.Println(RED + "Uh oh, there was an error getting your talking to GitHub. Here's what we got back:\n" + CLEAR)
		fmt.Println(string(body))
	}

	return body
}
