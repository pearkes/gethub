package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCloneRepo struct{}

func (*stepCloneRepo) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepCloneRepo) Cleanup(map[string]interface{}) {}
