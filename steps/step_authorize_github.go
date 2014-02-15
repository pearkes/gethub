package steps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/mitchellh/multistep"
)

type StepAuthorizeGithub struct{}

type AuthorizeResponse struct {
	Token string `json:"token"`
}

// The authorization sequence, required for someone without a ~/.getconfig
func (*StepAuthorizeGithub) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Begin authorization sequence...")

	pwd := os.Getenv("PWD")
	fmt.Printf("Enter the path to store respositories (%s if blank): ", pwd)

	var path string
	_, err := fmt.Scanf("%s", &path)

	if path == "" {
		path = pwd
	}

	// Let the user know what will happen.
	fmt.Println(`
Your username and password will be used once to obtain a unique
authorization token from GitHub's API, which will be stored in
~/.gethubconfig.
`)

	// Asks the user for credentials, and then makes a request to the
	// GitHub API to get an authorization token to store in ~/.getconfig
	log.Println("Asking user for credentials...")

	fmt.Printf("Please enter your GitHub username: ")
	var username string
	_, err = fmt.Scanf("%s", &username)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Please enter your GitHub password: ")

	p := gopass.GetPasswd()
	fmt.Printf("\n")

	password := string(p)

	client := &http.Client{}

	reqBody := strings.NewReader(`
        {"scopes":["repo"],
        "note":"gethub command line client",
        "note_url": "https://github.com/pearkes/gethub"}`)

	req, err := http.NewRequest("POST",
		"https://api.github.com/authorizations", reqBody)

	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	dec := json.NewDecoder(bytes.NewReader(body))

	log.Println(dec)

	auth := AuthorizeResponse{}

	dec.Decode(&auth)

	log.Println(resp.Status)

	if resp.StatusCode != 201 {
		fmt.Println(RED + "Uh oh, there was an error authenticating with GitHub. Here's what we got back:\n" + CLEAR)
		fmt.Println(string(body))
		return multistep.ActionHalt
	} else {
		fmt.Println(GREEN + "Succesfully authenticated with Github. Try running `gethub`." + CLEAR)
	}

	log.Println(string(body))

	// Set the discovered credentials into the bag of state
	state.Put("token", auth.Token)
	state.Put("username", username)
	state.Put("path", path)

	log.Println(username)

	return multistep.ActionContinue
}

func (*StepAuthorizeGithub) Cleanup(multistep.StateBag) {}
