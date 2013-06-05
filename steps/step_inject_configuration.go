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
	c, _ := config.ReadDefault(os.Getenv("HOME") + "/.getconfig")

	ignoredRepos, _ := c.String("ignores", "repo")
	ignoredOwners, _ := c.String("ignores", "owner")

	state["path"], _ = c.String("get", "path")
	state["token"], _ = c.String("github", "token")
	state["username"], _ = c.String("github", "username")
	state["ignored_repos"] = strings.Split(ignoredRepos, ",")
	state["ignored_owners"] = strings.Split(ignoredOwners, ",")

	return multistep.ActionContinue
}

func (*StepInjectConfiguration) Cleanup(map[string]interface{}) {}
