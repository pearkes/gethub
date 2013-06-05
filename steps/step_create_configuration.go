package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCheckPath struct{}

func (*stepCreateConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepCreateConfiguration) Cleanup(map[string]interface{}) {}
