package steps

import (
	"log"
	"os"
	"strings"

	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
)

const githubAPIHost = "https://api.github.com"

type StepInjectConfiguration struct{}

func (*StepInjectConfiguration) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Injecting configuration...")

	configPath := state.Get("config_path").(string)

	var path string

	// Determine if we are dealing with a custom config path
	if configPath == "" {
		// Default to the home directory
		path = os.Getenv("HOME") + "/.gethubconfig"
	} else {
		// They've specified a custom config path
		log.Println("Environment specified config path", configPath)
		path = configPath + ".gethubconfig"
	}

	// Read the file
	c, err := config.ReadDefault(path)

	log.Println(err)

	ignoredReposDirty, _ := c.String("ignores", "repo")
	ignoredOwnersDirty, _ := c.String("ignores", "owner")

	owners := []string{}
	repos := []string{}

	// Trim whitespace from the user configuration
	for _, ignoredRepo := range strings.Split(ignoredReposDirty, ",") {
		repos = append(repos, strings.TrimSpace(ignoredRepo))
	}
	for _, ignoredOwner := range strings.Split(ignoredOwnersDirty, ",") {
		owners = append(owners, strings.TrimSpace(ignoredOwner))
	}

	gpath, _ := c.String("gethub", "path")
	state.Put("path", gpath)
	token, _ := c.String("github", "token")
	state.Put("token", token)
	username, _ := c.String("github", "username")
	state.Put("username", username)
	host, _ := c.String("github", "host")
	if host == "" {
		host = githubAPIHost
	}
	state.Put("host", host)

	state.Put("ignored_repos", repos)
	state.Put("ignored_owners", owners)

	return multistep.ActionContinue
}

func (*StepInjectConfiguration) Cleanup(multistep.StateBag) {}
