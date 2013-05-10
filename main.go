package main

import (
	"bytes"
	"encoding/json"
	"flag"
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
func askForCredentials(env Env) string {
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
		fmt.Println("\x1b[1;31;40mUh oh, there was an error authenticating with GitHub. Here's what we got back:\x1b[0m\n")
		fmt.Println(string(body))
		os.Exit(1)
	} else {
		fmt.Println("\x1b[32mSuccesfully authenticated with Github.\x1b[0m")
	}

	log.Println(string(body))

	return auth.Token
}

// Parses options sent to `get` and kicks off the main event.
func main() {
	// Debugging and version flags
	debug := flag.Bool("debug", false, "Logs debugging information to STDOUT.")
	version := flag.Bool("version", false, "Prints the version and exits.")
	flag.BoolVar(debug, "d", false, "Logs debugging information to STDOUT.")
	flag.BoolVar(version, "v", false, "Prints the version and exits.")

	// Override the flag.Usage function to print custom usage info.
	flag.Usage = usage
	flag.Parse()

	// Discard logging if debug is turned off.
	if *debug == false {
		log.SetOutput(ioutil.Discard)
	}

	// Print the version and exit
	if *version {
		fmt.Println(versionString())
		os.Exit(1)
	}

	// Log enabled debugging
	log.Println("Debugging enabled for", versionString())

	env := Env{ProvidedPath: flag.Arg(0), Debug: *debug}

	// Run checks
	env = sequence_checks(env)

	// Update reposotories
	sequence_update(env)
}
