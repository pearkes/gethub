package main

import (
	"fmt"
	"github.com/pearkes/goconfig/config"
	"log"
	"os"
	"strings"
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
	Username      string
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

// Creates configuration at ~/.getconfig.
func createConfiguration() {
	log.Println("Creating configuration...")
}

// Injects the configuration into the environment.
func injectConfiguration() Configuration {
	log.Println("Injecting configuration...")

	// Read the file from their home directory
	c, err := config.ReadDefault(os.Getenv("HOME") + "/.getconfig")

	if err != nil {
		fmt.Println("Error reading from ~/.getconfig.")
	}

	path, _ := c.String("get", "path")
	token, _ := c.String("github", "token")
	username, _ := c.String("github", "username")
	repoIgnores, _ := c.String("ignores", "repo")
	ownerIgnores, _ := c.String("ignores", "owner")

	log.Println("Configured path:", path)
	log.Println("Configured username:", username)
	log.Println("Configured ignored repos:", repoIgnores)
	log.Println("Configured ignored owners:", ownerIgnores)

	conf := Configuration{
		Path:          path,
		Username:      username,
		Token:         token,
		IgnoredRepos:  strings.Split(repoIgnores, ","),
		IgnoredOwners: strings.Split(ownerIgnores, ","),
	}

	return conf
}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func checkConfiguration(env Env) {
	log.Println("Checking configuration...")

	// Check to see if the file exists at all. If not, drop into
	// the authorization sequence.
	_, err := os.Stat(os.Getenv("HOME") + "/.getconfig")

	if err != nil {
		sequence_authorize(env)
	}

	// Read the file from their home directory
	_, err2 := config.ReadDefault(os.Getenv("HOME") + "/.getconfig")

	if err2 != nil {
		fmt.Println("Your ~/.getconfig file may be corrupt. Try deleting it?")
	}

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
