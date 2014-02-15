package steps

import (
	"log"
	"os"

	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
)

type StepCreateConfiguration struct{}

func (*StepCreateConfiguration) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Creating configuration...")

	path := state.Get("path").(string)
	username := state.Get("username").(string)
	token := state.Get("token").(string)

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

func (*StepCreateConfiguration) Cleanup(multistep.StateBag) {}
