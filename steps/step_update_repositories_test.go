package steps

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mitchellh/multistep"
)

func TestStepUpdateRepositories(t *testing.T) {
	env := new(multistep.BasicStateBag)
	env.Put("path", "tmp")
	env.Put("ignored_owners", []string{})
	env.Put("ignored_repos", []string{})

	originPath := "tmp/pearkes/origin"

	os.MkdirAll(originPath, 0777)

	// Create a fake repository to clone from
	cmdInit := exec.Command("git", "init", originPath)

	// Commit to the repoistory to avoid warnings
	os.Create(originPath + "/test")
	cmdAdd := exec.Command("git", "add", "test")
	cmdAdd.Dir = originPath
	cmdCommit := exec.Command("git", "commit", "-m", "initial commit")
	cmdCommit.Dir = originPath

	// Clone the repository
	cmdClone := exec.Command("git", "clone", originPath, "tmp/pearkes/test")

	cmdInit.Run()
	cmdAdd.Run()
	cmdCommit.Run()
	cmdClone.Run()

	repo := Repo{FullName: "pearkes/test", SSHUrl: "../origin"}
	env.Put("repos", []Repo{repo})

	step := &StepUpdateRepositories{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}
}
