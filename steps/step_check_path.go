package steps

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"os"
)

type StepCheckPath struct{}

func (*StepCheckPath) Run(state map[string]interface{}) multistep.StepAction {
	repoPath := state["path"].(string)

	stat, _ := os.Stat(repoPath)

	if stat.IsDir() != true {
		// If the configured path isn't a directory, tell the user.
		fmt.Println(RED + "Your configured path (~/.getconfig) doesn't appear to be a directory." + CLEAR)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*StepCheckPath) Cleanup(map[string]interface{}) {}
