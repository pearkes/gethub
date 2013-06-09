package steps

import (
	"github.com/mitchellh/multistep"
	"os"
	"testing"
)

func TestStepCheckPath_Exists(t *testing.T) {
	env := make(map[string]interface{})

	env["path"] = "tmp/"
	os.Mkdir("tmp", 0777)

	step := &StepCheckPath{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
	os.RemoveAll("tmp")
}

func TestStepCheckPath_Not_Exists(t *testing.T) {
	env := make(map[string]interface{})

	env["path"] = "foobar/"

	step := &StepCheckPath{}

	results := step.Run(env)

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionContinue")
	}
}
