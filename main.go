package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Asks the user for credentials, and then makes a request to the
// GitHub API to get an authorization token to store in ~/.getconfig
func askForCredentials() {
	log.Println("Asking for credentials...")
}

// Retrieves a list of all available repostories and builds them up into
// something we can handle locally. After this occurs, we begin our
// clone / fetch sequence.
func listRepostories(env Env) {
	log.Println("Listing repostories...")
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

	conf := injectConfiguration()
	env := Env{ProvidedPath: flag.Arg(0), Debug: *debug, Config: conf}

	// Run checks
	env = sequence_checks(env)
	// Update reposotories
	sequence_update(env)
}
