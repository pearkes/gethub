package steps

import (
	"github.com/mitchellh/multistep"
	"log"
	"os"
)

type StepCheckRepo struct{}

func (*StepCheckRepo) Run(state map[string]interface{}) multistep.StepAction {
	repo := state["repo"].(Repo)
	path := state["path"].(string)
	ignoredRepos := state["ignored_repos"].([]string)
	ignoredOwners := state["ignored_owners"].([]string)

	log.Println("Starting check for:", repo.FullName)

	repoPath := path + "/" + repo.FullName

	// Check if the repo is ignored by it's name
	for _, ignoredName := range ignoredRepos {
		if ignoredName == repo.Name() {
			state["repo_state"] = "ignore"
			return multistep.ActionContinue
		}
	}

	// Check if the repo is ignored by it's owner
	for _, ignoredOwner := range ignoredOwners {
		if ignoredOwner == repo.Owner() {
			state["repo_state"] = "ignore"
			return multistep.ActionContinue
		}
	}

	// The path to the expected git internals
	gitPath := repoPath + "/.git"

	// Check to see if the directory is a git repository
	stat, staterr := os.Stat(gitPath)

	if staterr != nil || stat.IsDir() != true {
		state["repo_state"] = "clone"
		return multistep.ActionContinue
	} else {
		// If the directory does exist, we want to run a fetch on it.
		state["repo_state"] = "fetch"
		return multistep.ActionContinue
	}

	return multistep.ActionContinue
}

func (*StepCheckRepo) Cleanup(map[string]interface{}) {}
