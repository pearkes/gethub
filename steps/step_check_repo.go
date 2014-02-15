package steps

import (
	"log"
	"os"

	"github.com/mitchellh/multistep"
)

type StepCheckRepo struct{}

func (*StepCheckRepo) Run(state multistep.StateBag) multistep.StepAction {
	repo := state.Get("repo").(Repo)
	path := state.Get("path").(string)
	ignoredRepos := state.Get("ignored_repos").([]string)
	ignoredOwners := state.Get("ignored_owners").([]string)

	log.Println("Starting check for:", repo.FullName)

	repoPath := path + "/" + repo.FullName

	// Check if the repo is ignored by it's name
	for _, ignoredName := range ignoredRepos {
		if ignoredName == repo.Name() {
			state.Put("repo_state", "ignore")
			state.Put("repo_result", "ignore")
			return multistep.ActionContinue
		}
	}

	// Check if the repo is ignored by it's owner
	for _, ignoredOwner := range ignoredOwners {
		if ignoredOwner == repo.Owner() {
			state.Put("repo_state", "ignore")
			state.Put("repo_result", "ignore")
			return multistep.ActionContinue
		}
	}

	// The path to the expected git internals
	gitPath := repoPath + "/.git"

	// Check to see if the directory is a git repository
	stat, staterr := os.Stat(gitPath)

	if staterr != nil || stat.IsDir() != true {
		state.Put("repo_state", "clone")
		return multistep.ActionContinue
	} else {
		// If the directory does exist, we want to run a fetch on it.
		state.Put("repo_state", "fetch")
		return multistep.ActionContinue
	}

}

func (*StepCheckRepo) Cleanup(multistep.StateBag) {}
