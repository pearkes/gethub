package steps

import (
	"github.com/mitchellh/multistep"
)

type stepInjectConfiguration struct{}

func (*stepInjectConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Injecting configuration...")

	// Read the file from their home directory
	c, err := config.ReadDefault(os.Getenv("HOME") + "/.getconfig")

	if err != nil {
		fmt.Println("Error reading from ~/.getconfig")
		return multistep.ActionHalt
	}

	ignoredRepos, _ := c.String("ignores", "repo")
	ignoredOwners, _ := c.String("ignores", "owner")

	state["path"], _ = c.String("get", "path")
	state["token"], _ = c.String("github", "token")
	state["username"], _ = c.String("github", "username")
	state["ignored_repos"] = strings.Split(ignoredRepos, ",")
	state["ignored_owners"] = strings.Split(ignoredOwners, ",")

	return multistep.ActionContinue
}

func (*stepInjectConfiguration) Cleanup(map[string]interface{}) {}
