package main

import (
	"fmt"
	"log"
	"sync"
)

// The authorization sequence, required for someone without a ~/.getconfig
func sequence_authorize(env Env) Env {
	log.Println("Begin authorization sequence...")

	// Let the user know what will happen.
	fmt.Println(`Your username and password will be used once to obtain a unique
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

	fetches := []string{}
	clones := []string{}
	errors := []string{}
	ignores := []string{}

	// Asynchronously update each repository
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo Repo) {
			switch checkRepo(repo, env) {
			case "fetch":
				fetches = append(fetches, repo.Name())
			case "clone":
				clones = append(clones, repo.Name())
			case "error":
				errors = append(errors, repo.Name())
			case "ignore":
				ignores = append(ignores, repo.Name())
			}
			wg.Done()
		}(repo)
	}

	// Wait for every update to be finished
	wg.Wait()

	fmt.Println("Updated repositories:", len(fetches))

	fmt.Println("New repositories:", len(clones))
	for _, repo := range clones {
		fmt.Println("\t", repo)
	}

	if len(errors) > 0 {
		fmt.Println(len(errors), "error(s) encountered.")
	}
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
