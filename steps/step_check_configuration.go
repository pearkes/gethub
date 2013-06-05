package steps

import (
	"github.com/mitchellh/multistep"
)

type StepCheckConfiguration struct{}

func (*StepCheckConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*StepCheckConfiguration) Cleanup(map[string]interface{}) {}
