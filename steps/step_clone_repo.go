package steps

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mitchellh/multistep"
)

type StepCloneRepo struct{}

func (*StepCloneRepo) Run(state multistep.StateBag) multistep.StepAction {
	repoState := state.Get("repo_state").(string)

	if repoState != "clone" {
		log.Println("Skipping clone, repo state is " + repoState)
		return multistep.ActionContinue
	}

	repo := state.Get("repo").(Repo)
	path := state.Get("path").(string)

	repoPath := path + "/" + repo.FullName

	log.Println("Cloning new repository:", repoPath)

	// Make the repository directory
	mkdirerr := os.MkdirAll(repoPath, 0777)

	// If an error occurs, log and halt
	if mkdirerr != nil {
		log.Println("Error creating directory for " + repo.FullName)
		fmt.Printf("%s.%s", RED, CLEAR)
		state.Put("repo_result", "error")
		return multistep.ActionHalt
	}

	// Clone into the current directory
	cmd := exec.Command("git", "clone", repo.SSHUrl, ".")

	// Set the current directory as the path to the repository
	cmd.Dir = repoPath

	// Execute the clone
	cloneerr := cmd.Run()

	// If an error occurs, log and halt
	if cloneerr != nil {
		log.Println("Error cloning " + repo.FullName)
		fmt.Printf("%s.%s", RED, CLEAR)
		state.Put("repo_result", "error")
		return multistep.ActionHalt
	}

	// Print a success dot
	fmt.Printf("%s.%s", GREEN, CLEAR)
	state.Put("repo_result", "clone")
	return multistep.ActionContinue
}

func (*StepCloneRepo) Cleanup(multistep.StateBag) {}
