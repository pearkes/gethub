package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/multistep"
	"github.com/pearkes/gethub/steps"
)

func main() {
	// Debugging and version flags
	debug := flag.Bool("debug", false, "Logs debugging information to STDOUT.")
	version := flag.Bool("version", false, "Prints the version and exits.")
	flag.BoolVar(debug, "d", false, "Logs debugging information to STDOUT.")
	flag.BoolVar(version, "v", false, "Prints the version and exits.")

	// Override the flag.Usage function to print custom usage info.
	flag.Usage = usage
	flag.Parse()
	arg := flag.Arg(0)

	// Discard logging if debug is turned off.
	if *debug == false {
		log.SetOutput(ioutil.Discard)
	}

	// Print the version and exit
	if *version {
		fmt.Println(versionString())
		os.Exit(1)
	}

	// Log enabled debugging
	log.Println("Debugging enabled for", versionString())

	state := new(multistep.BasicStateBag)
	state.Put("debug", *debug)
	state.Put("config_path", os.Getenv("GETCONFIG_PATH"))

	if arg == "authorize" {
		authorizeRunner(state)
	} else if arg == "" {
		updateRunner(state)
	} else {
		fmt.Println("Invalid argument: " + arg)
		// Prints the usage and exits
		usage()
	}
}

// Builds the steps and kicks off the runner for updating
// repositories.
func updateRunner(state multistep.StateBag) {

	steps := []multistep.Step{
		&steps.StepCheckConfigurationFile{},
		&steps.StepInjectConfiguration{},
		&steps.StepCheckPath{},
		&steps.StepCheckConfiguration{},
		&steps.StepRetrieveRepositories{},
		&steps.StepUpdateRepositories{},
	}

	runner := &multistep.BasicRunner{Steps: steps}
	runner.Run(state)
}

// Builds the steps and kicks off the runner for authorizing
// and creating configuration.
func authorizeRunner(state multistep.StateBag) {

	steps := []multistep.Step{
		&steps.StepAuthorizeGithub{},
		&steps.StepCreateConfiguration{},
		&steps.StepCheckConfigurationFile{},
		&steps.StepCheckConfiguration{},
	}

	runner := &multistep.BasicRunner{Steps: steps}
	runner.Run(state)
}
