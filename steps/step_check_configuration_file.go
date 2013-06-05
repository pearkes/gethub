package steps

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
	"log"
	"os"
)

type StepCheckConfigurationFile struct{}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func (*StepCheckConfigurationFile) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Checking configuration...")

	var configPath string
	path, _ := state["path"].(string)

	// Determine if we are dealing with a custom config path
	if path == "" {
		// Default to the home directory
		configPath = os.Getenv("HOME") + "/.gethubconfig"
	} else {
		// They've specified a custom config path
		log.Println("Environment specified config path", path)
		configPath = path + "/.gethubconfig"
	}

	// Is the config file even there?
	_, err := os.Stat(configPath)

	if err != nil {
		fmt.Println(RED + "It seems as though you haven't set-up gethub. Please run `gethub authorize`" + CLEAR)
		return multistep.ActionHalt
	}

	// Read the file and see if all is well
	_, err2 := config.ReadDefault(configPath)

	if err2 != nil {
		fmt.Println(RED + "Something seems to be wrong with your ~/.gethubconfig file. Please run `gethub authorize`" + CLEAR)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*StepCheckConfigurationFile) Cleanup(map[string]interface{}) {}
