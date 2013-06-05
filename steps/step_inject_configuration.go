package steps

import (
	"github.com/mitchellh/multistep"
)

type stepInjectConfiguration struct{}

func (*stepInjectConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepInjectConfiguration) Cleanup(map[string]interface{}) {}
