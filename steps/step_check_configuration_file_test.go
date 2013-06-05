package steps

import (
	"github.com/mitchellh/multistep"
	"os"
	"testing"
)

func TestStepCheckConfigurationFile_No_Config(t *testing.T) {
	env := make(map[string]interface{})

	os.Mkdir("tmp", 0777)
	env["config_path"] = "tmp/"

	results := StepCheckConfigurationFile.Run(env)
	// Output: It seems as though you haven't set-up gethub. Please run `gethub authorize`

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionHalt")
	}
}

func TestStepCheckConfigurationFile_Corrupt_Config(t *testing.T) {
	env = make(map[string]interface{})

	os.Mkdir("tmp", "0777")
	env["config_path"] = "tmp/"
	file, _ := os.Create("tmp/.gethubconfig")

	// Some messy string
	file.WriteString("foobar:baz:bar\n\nfoob:ar")

	results := StepCheckConfigurationFile.Run(env)
	// Output: Something seems to be wrong with your ~/.gethubconfig file. Please run `gethub authorize`

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionHalt")
	}
}

func TestStepCheckConfigurationFile_Good_Config(t *testing.T) {
	env = make(map[string]interface{})

	os.Mkdir("tmp", "0777")
	env["config_path"] = "tmp/"
	file, _ := os.Create("tmp/.gethubconfig")

	results := StepCheckConfigurationFile.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
