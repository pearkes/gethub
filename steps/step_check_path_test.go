package steps

import (
	"github.com/mitchellh/multistep"
	"testing"
)

func TestStepCheckPath(t *testing.T) {
	env = make(map[string]interface{})

	results := stepCheckPath.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
