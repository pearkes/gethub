package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
)

func cloneRepo(repo Repo, env Env) error {
	repoPath := env.Config.Path + repo.FullName
	log.Println("Cloning new repository:", repoPath)

	// Clone into the current directory
	cmd := exec.Command("git", "clone", repo.SSHUrl, ".")

	// Set the current directory as the path to the repository
	cmd.Dir = repoPath

	// Grab stdout so we can log it if an error occurs
	var out bytes.Buffer
	cmd.Stdout = &out

	// Execute the clone
	err := cmd.Run()

	// If an error occurs, return a new error with the stdout
	if err != nil {
		return errors.New("Error cloning: " + out.String())
	}

	// If everything went well, return nil
	return nil
}

func fetchRepo(repo Repo, env Env) error {
	repoPath := env.Config.Path + repo.FullName

	log.Println("Fetching new repository:", repoPath)

	// Fetch the current directory
	cmd := exec.Command("git", "fetch")

	// Set the current directory as the path to the repository
	cmd.Dir = repoPath

	// Grab stdout so we can log it if an error occurs
	var out bytes.Buffer
	cmd.Stdout = &out

	// Execute the clone
	err := cmd.Run()

	// If an error occurs, return a new error with the stdout
	if err != nil {
		return errors.New("Error fetching: " + out.String())
	}

	// If everything went well, return nil
	return nil
}

// Checks a repostories in the get configured path.
// If the repository path doesn't exist, we should clone it.
// If the repository path does exist, we should run a fetch on it.
// All succesfull clones and fetches should be passed down their
// particular channel to communicate back to the user. Any errors
// encountered should be passed back through the error channel to
// notify the user of such.
// This returns "error", "fetch", "clone" or "ignore"
func checkRepo(repo Repo, env Env) string {
	log.Println("Starting check for:", repo.FullName)
	repoPath := env.Config.Path + repo.FullName

	// Check if the repo is ignored by it's name
	for _, ignoredName := range env.Config.IgnoredRepos {
		if ignoredName == repo.Name() {
			log.Println("Ignoring repository based on configuration:", repo.FullName)
			return "ignore"
		}
	}

	// Check if the repo is ignored by it's owner
	for _, ignoredOwner := range env.Config.IgnoredOwners {
		if ignoredOwner == repo.Owner() {
			log.Println("Ignoring repository based on configuration:", repo.FullName)
			return "ignore"
		}
	}

	// The path to the expected git internals
	gitPath := repoPath + "/.git"

	// Check to see if the directory is a git repository
	stat, _ := os.Stat(gitPath)

	if stat.IsDir() != true {
		// If the directory does not exist, we want to run a clone to
		// get it.
		cloneerr := cloneRepo(repo, env)

		if cloneerr != nil {
			// If there is a clone error, we should log it and return
			// that this repo failed with an error.
			log.Println(cloneerr)
			return "error"
		} else {
			// If there isn't a clone error, we should log that it was
			// succesful and return that it was a clone
			log.Println("Succesfully cloned repository:", repo.FullName)
			return "clone"
		}
	} else {
		// If the directory does exist, we want to run a fetch on it.
		fetcherr := fetchRepo(repo, env)

		if fetcherr != nil {
			// If there is a fetch error, we should log it and return
			// that this repo failed with an error.
			log.Println(fetcherr)
			return "error"
		} else {
			// If there isn't a fetch error, we should log that it was
			// succesful and return that it was a fetch
			log.Println("Succesfully fetched repository:", repo.FullName)
			return "fetch"
		}
	}

	return "unknown"
}
