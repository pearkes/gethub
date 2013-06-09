package steps

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"log"
	"os/exec"
)

type StepFetchRepo struct{}

func (*StepFetchRepo) Run(state map[string]interface{}) multistep.StepAction {
	repoState := state["repo_state"].(string)

	if repoState != "fetch" {
		log.Println("Skipping clone, repo state is " + repoState)
		return multistep.ActionContinue
	}

	repo := state["repo"].(Repo)
	path := state["path"].(string)

	repoPath := path + "/" + repo.FullName

	log.Println("Fetching existing repository:", repoPath)

	// Fetch the current directory
	cmd := exec.Command("git", "fetch")

	// Set the current directory as the path to the repository
	cmd.Dir = repoPath

	// Execute the clone
	err := cmd.Run()

	// If an error occurs, return a new error with the stdout
	if err != nil {
		log.Println("Error fetching " + repo.FullName)
		fmt.Printf("%s.%s", RED, CLEAR)
		state["repo_result"] = "error"
		return multistep.ActionHalt
	}

	// Print a success dot
	fmt.Printf("%s.%s", GREEN, CLEAR)
	state["repo_result"] = "fetch"
	return multistep.ActionContinue
}

func (*StepFetchRepo) Cleanup(map[string]interface{}) {}
