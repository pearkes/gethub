package main

import (
	"fmt"
	"log"
)

// The authorization sequence, required for someone without a ~/.getconfig
func sequence_authorize(env Env) Env {
	log.Println("Begin authorization sequence...")

	// Let the user know what will happen.
	fmt.Println(`
Your username and password will be used once to obtain a unique
authorization token from GitHub's API, which will be stored in
~/.getconfig.
`)

	env = askForCredentials(env)

	log.Println(env.Config.Username, env.Config.Token)

	// Make the configuration file.
	createConfiguration(env)

	return env
}

// The update sequence, which retrieves the repos and fetches or clones
// each of them.
func sequence_update(env Env) {
	log.Println("Begin repository update sequence...")

	repos := listRemoteRepostories(env)

	checkRepo(repos[1], env)
}

// The check sequence, which goes through the basic health checks for
// `get` to succesfully function.
func sequence_checks(env Env) Env {
	log.Println("Begin check sequence...")

	// Inject configuration
	env.Config = injectConfiguration()

	// Check supplied path
	checkPath(env)

	// Check Configuration
	checkConfiguration(env)

	// Check configured path
	checkPath(env)

	return env
}
