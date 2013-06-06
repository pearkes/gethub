package steps

import (
	"github.com/mitchellh/multistep"
	"os"
	"os/exec"
	"testing"
)

func TestStepCloneRepo(t *testing.T) {
	env := make(map[string]interface{})
	env["path"] = "tmp"
	os.MkdirAll("tmp/pearkes/origin", 0777)

	// Create a fake repository to clone from
	cmdInit := exec.Command("git", "init", "tmp/pearkes/origin")
	// Commit to the repoistory to avoid warnings
	os.Create("tmp/pearkes/origin/test")
	cmdCommit := exec.Command("git", "commit", "-a", "-m", "'initial commit'")

	cmdInit.Run()
	cmdCommit.Run()

	repo := Repo{FullName: "pearkes/test", SSHUrl: "tmp/pearkes/origin"}
	env["repo"] = repo
	env["repo_state"] = "clone"

	step := &StepCloneRepo{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
