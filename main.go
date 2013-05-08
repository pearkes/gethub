package main

import (
	"flag"
	"fmt"
	"os"
)

// We pass around the environment and attach various useful things to
// it, like arguments to `get` and the configuration.
type Env struct {
	Config       Configuration
	ProvidedPath string
	Path         string
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

}

// Injects the configuration into the environment.
func injectConfiguration() {
	// conf := Configiration{}
}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func checkConfiguration(conf Configuration) {

}

// Asks the user for credentials, and then makes a request to the
// GitHub API to get an authorization token to store in ~/.getconfig
func askForCredentials() {

}

// Checks a path to see if it is get compatible. If not, it raises an
// error.
func checkPath(env Env) {
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

}

// The authorization sequence, required for someone without a ~/.getconfig
func sequence_authorize(env Env) Env {
	return env
}

// The update sequence, which retrieves the repos and fetches or clones
// each of them.
func sequence_update() {

}

// The check sequence, which goes through the basic health checks for
// `get` to succesfully function.
func sequence_checks(env Env) Env {
	checkConfiguration(env.Config)
	checkPath(env)
	return env
}

// Parses options sent to `get` and kicks off the main event.
func main() {
	flag.Parse()

	// conf := injectConfiguration()

	env := Env{ProvidedPath: flag.Arg(0)}

	// Run checks
	env = sequence_checks(env)
	//
	{

	}
}

func usage() {
	fmt.Println(`Usage: get [<path>] [-v] [-h] [-d]

    -v, --version                   Prints the version and exits.
    -h, --help                      Prints the usage information.
    -d, --debug                     Logs debugging information to STDOUT.

Arguments:

    path                            The path to place or update the
                                    repostories. Defaults to the path
                                    in ~/.get.

To learn more or to contribute, please see github.com/pearkes/get`)
	os.Exit(1)
}
