package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCheckRepo struct{}

func (*stepCheckRepo) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepCheckRepo) Cleanup(map[string]interface{}) {}
