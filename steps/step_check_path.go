package steps

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"log"
	"os"
)

type StepCheckPath struct{}

func (*StepCheckPath) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Checking path...")

	repoPath := state["path"].(string)

	stat, err := os.Stat(repoPath)

	if err != nil || stat.IsDir() != true {
		// If the configured path isn't a directory, tell the user.
		fmt.Println(RED + "Your configured path (~/.gethubconfig) doesn't appear to be a directory." + CLEAR)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*StepCheckPath) Cleanup(map[string]interface{}) {}
