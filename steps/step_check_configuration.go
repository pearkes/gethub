package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCheckConfiguration struct{}

func (*stepCheckConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepCheckConfiguration) Cleanup(map[string]interface{}) {}
