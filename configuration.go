package main

import (
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
func injectConfiguration() Configuration {
	log.Println("Injecting configuration...")
	conf := Configuration{}
	return conf
}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func checkConfiguration(conf Configuration) {
	log.Println("Checking configuration...")
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
