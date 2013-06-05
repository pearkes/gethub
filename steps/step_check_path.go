package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCheckPath struct{}

func (*stepCheckPath) Run(state map[string]interface{}) multistep.StepAction {
	stat, _ := os.Stat(state["path"])

	if stat.IsDir() != true {
		// If the configured path isn't a directory, tell the user.
		fmt.Println(RED + "Your configured path (~/.getconfig) doesn't appear to be a directory." + CLEAR)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*stepCheckPath) Cleanup(map[string]interface{}) {}
