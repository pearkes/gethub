package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCheckPath struct{}

func (*stepCheckPath) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepCheckPath) Cleanup(map[string]interface{}) {}
