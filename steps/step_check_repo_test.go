package steps

import (
	"github.com/mitchellh/multistep"
	"testing"
)

func TestStepCheckRepo(t *testing.T) {
	env = make(map[string]interface{})

	results := stepCheckRepo.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
