package main

import (
	"fmt"
	"log"
	"strings"
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

	fmt.Printf("Contacting GitHub... ")

	repos := listRemoteRepostories(env)

	fmt.Printf("%sdone%s\n", green, clear)

	fmt.Printf("Updating repositories: ")

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
				fmt.Printf("%s.%s", green, clear)
			case "clone":
				clones = append(clones, repo.Name())
				fmt.Printf("%s.%s", green, clear)
			case "error":
				errors = append(errors, repo.Name())
				fmt.Printf("%s.%s", red, clear)
			case "ignore":
				ignores = append(ignores, repo.Name())
			}
			wg.Done()
		}(repo)
	}

	// Wait for every update to be finished
	wg.Wait()

	mess := []string{}

	if len(fetches) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d repos updated%s", green, len(fetches), clear))
	}

	if len(clones) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d new repos%s (%s)", green, len(clones), clear, strings.Join(clones, ", ")))
	}

	if len(errors) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d errors%s", red, len(errors), clear))
	}

	fmt.Printf("\n%s\n", strings.Join(mess, ", "))
}

// The check sequence, which goes through the basic health checks for
// `get` to succesfully function.
func sequence_checks(env Env) Env {
	log.Println("Begin check sequence...")

	// Check Configuration
	checkConfiguration(env)

	// Inject configuration
	env.Config = injectConfiguration()

	// Check path
	checkPath(env)

	return env
}
