package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// We pass around the environment and attach various useful things to
// it, like arguments to `get` and the configuration.
type Env struct {
	Config       Configuration
	ProvidedPath string
	Path         string
	Debug        bool
}

// type Configuration is built from the file ~/.getconfig.
type Configuration struct {
	Token         string
	Path          string
	IgnoredOwners []string
	IgnoredRepos  []string
}

// type Locals represents the list of local owners (1st level) and
// local repos (2nd level). It is read from the file system, from the
// path specified by either an argument or the configuration.
type Locals struct {
	Owners []string
	Repos  []string
}

// type Remotes represents the list of remote owners and repos for
// each of those retrieved from the GitHub API.
type Remotes struct {
	Owners []string
	Repos  []string
}

// Creates configuration at ~/.getconfig.
func createConfiguration() {
	log.Println("Creating configuration...")

}

// Injects the configuration into the environment.
func injectConfiguration() {
	log.Println("Injecting configuration...")

	// conf := Configiration{}
}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func checkConfiguration(conf Configuration) {
	log.Println("Checking configuration...")

}

// Asks the user for credentials, and then makes a request to the
// GitHub API to get an authorization token to store in ~/.getconfig
func askForCredentials() {
	log.Println("Asking for credentials...")

}

// Checks a path to see if it is get compatible. If not, it raises an
// error.
func checkPath(env Env) {
	log.Println("Checking path...")

	var path string

	if env.ProvidedPath != "" {
		path = env.ProvidedPath

	} else if env.Config.Path != "" {
		path = env.Config.Path
	}

	_, err := os.Stat(path)
	if err != nil {
		// They haven't set-up a path, let's take them through it.
		sequence_authorize(env)
	}
}

// Retrieves a list of all available repostories and builds them up into
// something we can handle locally. After this occurs, we begin our
// clone / fetch sequence.
func listRepostories() {
	log.Println("Listing repostories...")

}

// The authorization sequence, required for someone without a ~/.getconfig
func sequence_authorize(env Env) Env {
	log.Println("Begin authorization sequence...")

	return env
}

// The update sequence, which retrieves the repos and fetches or clones
// each of them.
func sequence_update() {
	log.Println("Begin repository update sequence...")

}

// The check sequence, which goes through the basic health checks for
// `get` to succesfully function.
func sequence_checks(env Env) Env {
	log.Println("Begin check sequence...")
	checkConfiguration(env.Config)
	checkPath(env)
	return env
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

	// conf := injectConfiguration()

	env := Env{ProvidedPath: flag.Arg(0), Debug: *debug}

	// Run checks
	env = sequence_checks(env)
	//
	{

	}
}
