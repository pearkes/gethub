package steps

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/multistep"
	"github.com/pearkes/goconfig/config"
)

type StepCheckConfigurationFile struct{}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func (*StepCheckConfigurationFile) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Checking configuration file...")
	configPath := state.Get("config_path").(string)
	var path string

	// Determine if we are dealing with a custom config path
	if configPath == "" {
		// Default to the home directory
		path = os.Getenv("HOME") + "/.gethubconfig"
	} else {
		// They've specified a custom config path
		log.Println("Environment specified config path", configPath)
		path = configPath + "/.gethubconfig"
	}

	// Is the config file even there?
	_, err := os.Stat(path)

	if err != nil {
		fmt.Println(RED + "It seems as though you haven't set-up gethub. Please run `gethub authorize`" + CLEAR)
		return multistep.ActionHalt
	}

	// Read the file and see if all is well with a basic config
	c, err2 := config.ReadDefault(path)
	checkPath, _ := c.String("gethub", "path")

	if checkPath == "" || err2 != nil {
		fmt.Println(RED + "Something seems to be wrong with your ~/.gethubconfig file. Please run `gethub authorize`" + CLEAR)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*StepCheckConfigurationFile) Cleanup(multistep.StateBag) {}
