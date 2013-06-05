package steps

import (
	"github.com/mitchellh/multistep"
)

type stepFetchRepo struct{}

func (*stepFetchRepo) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepFetchRepo) Cleanup(map[string]interface{}) {}
