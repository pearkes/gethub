package steps

import (
	"log"

	"github.com/mitchellh/multistep"
)

type StepCheckConfiguration struct{}

func (*StepCheckConfiguration) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Checking configuration...")
	// TODO Check the configuration.
	return multistep.ActionContinue
}

func (*StepCheckConfiguration) Cleanup(multistep.StateBag) {}
