package steps

import (
	"os"
	"testing"

	"github.com/mitchellh/multistep"
)

func TestStepCheckConfigurationFile_No_Config(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.Mkdir("tmp", 0777)
	env.Put("config_path", "tmp/")

	step := &StepCheckConfigurationFile{}

	results := step.Run(env)
	// Output: It seems as though you haven't set-up gethub. Please run `gethub authorize`

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionHalt")
	}
	os.RemoveAll("tmp")
}

func TestStepCheckConfigurationFile_Corrupt_Config(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.Mkdir("tmp", 0777)

	env.Put("config_path", "tmp/")
	file, _ := os.Create("tmp/.gethubconfig")

	// Some messy string
	file.WriteString("foobar!baz:bar\n\nfoob:ar")

	step := &StepCheckConfigurationFile{}

	results := step.Run(env)
	// Output: Something seems to be wrong with your ~/.gethubconfig file. Please run `gethub authorize`

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionHalt")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckConfigurationFile_Good_Config(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.Mkdir("tmp", 0777)
	env.Put("config_path", "tmp/")

	file, _ := os.Create("tmp/.gethubconfig")

	file.WriteString("\n[gethub]\npath: tmp/\n")

	step := &StepCheckConfigurationFile{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
