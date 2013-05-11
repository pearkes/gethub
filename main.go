package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// How do you make constants properly? Will accept pull requests.
var (
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	clear  = "\x1b[0m"
)

// Parses options sent to `get` and kicks off the main event.
func main() {
	// Debugging and version flags
	debug := flag.Bool("debug", false, "Logs debugging information to STDOUT.")
	version := flag.Bool("version", false, "Prints the version and exits.")
	flag.BoolVar(debug, "d", false, "Logs debugging information to STDOUT.")
	flag.BoolVar(version, "v", false, "Prints the version and exits.")

	// Override the flag.Usage function to print custom usage info.
	flag.Usage = usage
	flag.Parse()

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

	env := Env{ProvidedPath: flag.Arg(0), Debug: *debug}

	// Run checks
	env = sequence_checks(env)

	// Update reposotories
	sequence_update(env)
}
