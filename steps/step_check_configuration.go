package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCheckConfiguration struct{}

// Checks the configuration on the filesystem for syntax errors or
// non-exsistance.
func (*stepCheckConfiguration) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Checking configuration...")

	var configPath string

	// Determine if we are dealing with a custom config path
	if state["config_path"] == "" {
		// Default to the home directory
		configPath = os.Getenv("HOME") + "/.getconfig"
	} else {
		// They've specified a custom config path
		log.Println("Environment specified config path " + state["config_path"])
		configPath = state["config_path"] + "/.getconfig"
	}

	// Is the config file even there?
	_, err := os.Stat(configPath)

	if err != nil {
		fmt.Println(red + "It seems as though you haven't set-up get. Please run `get authorize`" + clear)
		return multistep.ActionHalt
	}

	// Read the file and see if all is well
	_, err2 := config.ReadDefault(configPath)

	if err2 != nil {
		fmt.Println(red + "Something seems to be wrong with your ~/.getconfig file. Please run `get authorize`" + clear)
		return multistep.ActionHalt
	}

	return multistep.ActionContinue
}

func (*stepCheckConfiguration) Cleanup(map[string]interface{}) {}
