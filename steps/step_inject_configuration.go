package steps

import (
	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
	"log"
	"os"
	"strings"
)

type StepInjectConfiguration struct{}

func (*StepInjectConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Injecting configuration...")

	// Read the file from their home directory
	c, _ := config.ReadDefault(os.Getenv("HOME") + "/.gethubconfig")

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

	state["path"], _ = c.String("gethub", "path")
	state["token"], _ = c.String("github", "token")
	state["username"], _ = c.String("github", "username")

	state["ignored_repos"] = repos
	state["ignored_owners"] = owners

	return multistep.ActionContinue
}

func (*StepInjectConfiguration) Cleanup(map[string]interface{}) {}
