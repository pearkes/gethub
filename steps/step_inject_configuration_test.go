package steps

import (
	"os"
	"testing"

	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
)

func TestStepInjectConfiguration(t *testing.T) {
	env := new(multistep.BasicStateBag)
	os.Mkdir("tmp", 0777)
	env.Put("config_path", "tmp/")

	conf := config.NewDefault()

	// Create the configuration file sections and items
	conf.AddSection("gethub")
	conf.AddSection("github")
	conf.AddSection("ignores")
	conf.AddOption("gethub", "path", "tmp")
	conf.AddOption("github", "username", "foo")
	conf.AddOption("github", "token", "bar")
	conf.AddOption("ignores", "repo", "facebook")
	conf.AddOption("ignores", "repo", "pearkes/bootstrap")
	conf.AddOption("ignores", "owner", "pearkes")

	conf.WriteFile("tmp/.gethubconfig", 0644, "")

	step := &StepInjectConfiguration{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
