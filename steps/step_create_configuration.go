package steps

import (
	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
	"log"
)

type stepCheckPath struct{}

func (*stepCreateConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Creating configuration...")
	conf := config.NewDefault()

	// Create the configuration file sections and items
	conf.AddSection("get")
	conf.AddSection("github")
	conf.AddSection("ignores")
	conf.AddOption("get", "path", state["provided_path"])
	conf.AddOption("github", "username", state["username"])
	conf.AddOption("github", "token", state["token"])
	conf.AddOption("ignores", "repo", "")
	conf.AddOption("ignores", "owner", "")

	conf.WriteFile(os.Getenv("HOME")+"/.getconfig", 0644, "")
	return multistep.ActionContinue
}

func (*stepCreateConfiguration) Cleanup(map[string]interface{}) {}
