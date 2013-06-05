package steps

import (
	"github.com/mitchellh/multistep"
)

type stepCloneRepo struct{}

func (*stepCloneRepo) Run(state map[string]interface{}) multistep.StepAction {
	if state["repo_state"] != "clone" {
		log.Println("Skipping clone, repo state is " + repo_state)
		return multistep.ActionContinue
	}

	repo := state["repo"]

	repoPath := state["repos_path"] + "/" + repo.FullName

	log.Println("Cloning new repository:", repoPath)

	// Make the repository directory
	mkdirerr := os.MkdirAll(repoPath, 0777)

	// If an error occurs, log and halt
	if mkdirerr != nil {
		log.Println(mkdirerr)
		fmt.Printf("%s.%s", RED, CLEAR)
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
		return multistep.ActionHalt
	}

	// Print a success dot
	fmt.Printf("%s.%s", GREEN, CLEAR)
	return multistep.ActionContinue
}

func (*stepCloneRepo) Cleanup(map[string]interface{}) {}
