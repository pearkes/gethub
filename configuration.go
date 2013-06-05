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

// Creates configuration at ~/.getconfig.
func createConfiguration(env Env) {
	log.Println("Creating configuration...")
	conf := config.NewDefault()

	// Create the configuration file sections and items
	conf.AddSection("get")
	conf.AddSection("github")
	conf.AddSection("ignores")
	conf.AddOption("get", "path", env.ProvidedPath)
	conf.AddOption("github", "username", env.Config.Username)
	conf.AddOption("github", "token", env.Config.Token)
	conf.AddOption("ignores", "repo", "")
	conf.AddOption("ignores", "owner", "")

	conf.WriteFile(os.Getenv("HOME")+"/.getconfig", 0644, "")
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

// Checks a path to see if it is get compatible. If not, it raises an
// error.
func stepCheckPath(env Env) {
	log.Println("Checking path...")

	if env.Config.Path != "" {
		// If we have the configuration, check the path provided there.
		stat, _ := os.Stat(env.Config.Path)

		if stat.IsDir() != true {
			// If the configured path isn't a directory, tell the user.
			fmt.Println(red + "Your configured path (~/.getconfig) doesn't appear to be a directory." + clear)
			os.Exit(1)
		}

	} else {
		// If we don't have configuration, perform a check on the provided
		// path.
		var path string

		if env.ProvidedPath != "" {
			path = env.ProvidedPath
		} else {
			fmt.Println(red + "You need to provide a path to clone your repositories to the first time your run get." + clear)
			usage()
		}

		_, err := os.Stat(path)

		if err != nil {
			// They haven't set-up a path, or passed one in, so we're going
			// to assume they want to do it here.
			fmt.Println(red + "Your provided path doesn't seem to exist: " + env.ProvidedPath + clear)
			os.Exit(1)
		}
	}

}
