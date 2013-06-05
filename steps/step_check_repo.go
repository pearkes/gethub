package steps

import (
	"github.com/mitchellh/multistep"
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
)

type stepCheckRepo struct{}

func (*stepCheckRepo) Run(state map[string]interface{}) multistep.StepAction {
    repo := state["repo"]

	log.Println("Starting check for:", repo.FullName)
	repoPath := state["repos_path"] + "/" + repo.FullName

	// Check if the repo is ignored by it's name
	for _, ignoredName := range state["ignored_repos"] {
		if ignoredName == repo.Name() {
			state["repo_state"] = "ignore"
		}
	}

	// Check if the repo is ignored by it's owner
	for _, ignoredOwner := state["ignored_owners"] {
		if ignoredOwner == repo.Owner() {
			state["repo_state"] = "ignore"
		}
	}

	// The path to the expected git internals
	gitPath := repoPath + "/.git"

	// Check to see if the directory is a git repository
	stat, staterr := os.Stat(gitPath)

	if staterr != nil || stat.IsDir() != true {
        state["repo_state"] = "clone"
	} else {
		// If the directory does exist, we want to run a fetch on it.
		state["repo_state"] = "fetch"
	}

	return multistep.ActionContinue
}

func (*stepCheckRepo) Cleanup(map[string]interface{}) {}
