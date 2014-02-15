package steps

import (
	"testing"

	"github.com/mitchellh/multistep"
)

func TestStepCheckConfiguration(t *testing.T) {

	env := new(multistep.BasicStateBag)

	step := &StepCheckConfiguration{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
