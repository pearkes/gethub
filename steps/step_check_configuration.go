package steps

import (
	"github.com/mitchellh/multistep"
	"log"
)

type StepCheckConfiguration struct{}

func (*StepCheckConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Checking configuration...")
	// Do some stuff
	return multistep.ActionContinue
}

func (*StepCheckConfiguration) Cleanup(map[string]interface{}) {}
