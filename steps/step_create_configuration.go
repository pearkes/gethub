package steps

import (
	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
	"log"
	"os"
)

type StepCreateConfiguration struct{}

func (*StepCreateConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Creating configuration...")

	path := state["path"].(string)
	username := state["username"].(string)
	token := state["token"].(string)

	conf := config.NewDefault()

	// Create the configuration file sections and items
	conf.AddSection("gethub")
	conf.AddSection("github")
	conf.AddSection("ignores")
	conf.AddOption("gethub", "path", path)
	conf.AddOption("github", "username", username)
	conf.AddOption("github", "token", token)
	conf.AddOption("ignores", "repo", "")
	conf.AddOption("ignores", "owner", "")

	conf.WriteFile(os.Getenv("HOME")+"/.gethubconfig", 0644, "")

	return multistep.ActionContinue
}

func (*StepCreateConfiguration) Cleanup(map[string]interface{}) {}
