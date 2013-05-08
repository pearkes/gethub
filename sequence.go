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

	askForCredentials()
	return env
}

// The update sequence, which retrieves the repos and fetches or clones
// each of them.
func sequence_update(env Env) {
	log.Println("Begin repository update sequence...")
	listRepostories(env)
}

// The check sequence, which goes through the basic health checks for
// `get` to succesfully function.
func sequence_checks(env Env) Env {
	log.Println("Begin check sequence...")
	checkConfiguration(env.Config)
	checkPath(env)
	return env
}
