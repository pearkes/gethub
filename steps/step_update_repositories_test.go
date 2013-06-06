package steps

import (
	"github.com/mitchellh/multistep"
	"testing"
)

func TestStepUpdateRepositories(t *testing.T) {
	env := make(map[string]interface{})

	step := &StepUpdateRepositories{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
