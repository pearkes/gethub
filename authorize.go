package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/howeyc/gopass"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type AuthorizeResponse struct {
	Token string `json:token`
}

// Asks the user for credentials, and then makes a request to the
// GitHub API to get an authorization token to store in ~/.getconfig
func askForCredentials(env Env) Env {
	log.Println("Asking user for credentials...")

	var username string

	fmt.Printf("Please enter your GitHub username: ")
	_, err := fmt.Scanf("%s", &username)

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
        "note":"get command line client",
        "note_url": "https://github.com/pearkes/get"}`)

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
		fmt.Println(red + "Uh oh, there was an error authenticating with GitHub. Here's what we got back:\n" + clear)
		fmt.Println(string(body))
		os.Exit(1)
	} else {
		fmt.Println(green + "Succesfully authenticated with Github." + clear)
	}

	log.Println(string(body))

	// Set the discovered credentials into the environment
	env.Config.Token = auth.Token
	env.Config.Username = username

	return env
}
