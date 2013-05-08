package main

import (
	"log"
)

// The authorization sequence, required for someone without a ~/.getconfig
func sequence_authorize(env Env) Env {
	log.Println("Begin authorization sequence...")

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
