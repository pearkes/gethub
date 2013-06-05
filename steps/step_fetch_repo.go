package steps

import (
	"github.com/mitchellh/multistep"
)

type stepFetchRepo struct{}

func (*stepFetchRepo) Run(state map[string]interface{}) multistep.StepAction {
	if state["repo_state"] != "fetch" {
		log.Println("Skipping clone, repo state is " + repo_state)
		return multistep.ActionContinue
	}

	repo := state["repo"]

	repoPath := state["repos_path"] + "/" + repo.FullName

	log.Println("Fetching existing repository:", repoPath)

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
		log.Println("Error fetching " + repo.FullName())
		fmt.Printf("%s.%s", RED, CLEAR)
		return multistep.ActionHalt
	}

	// Print a success dot
	fmt.Printf("%s.%s", GREEN, CLEAR)
	return multistep.ActionContinue
}

func (*stepFetchRepo) Cleanup(map[string]interface{}) {}
