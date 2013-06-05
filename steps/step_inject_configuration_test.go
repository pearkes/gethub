package steps

import (
	"github.com/mitchellh/multistep"
	"testing"
)

func TestStepInjectConfiguration(t *testing.T) {
	env = make(map[string]interface{})

	results := stepInjectConfiguration.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
