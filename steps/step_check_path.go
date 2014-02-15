package steps

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/multistep"
)

type StepCheckPath struct{}

func (*StepCheckPath) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Checking path...")

	repoPath := state.Get("path").(string)

	stat, err := os.Stat(repoPath)

	if err != nil || stat.IsDir() != true {
		// If the configured path isn't a directory, tell the user.
		fmt.Println(RED + "Your configured path (in ~/.gethubconfig) doesn't appear to be a directory." + CLEAR)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*StepCheckPath) Cleanup(multistep.StateBag) {}
