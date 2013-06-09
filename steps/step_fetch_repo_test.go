package steps

import (
	"github.com/mitchellh/multistep"
	"os"
	"os/exec"
	"testing"
)

func TestStepFetchRepo_Exists(t *testing.T) {
	env := make(map[string]interface{})
	env["path"] = "tmp"

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
	env["repo"] = repo
	env["repo_state"] = "fetch"

	step := &StepFetchRepo{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepFetchRepo_Not_Exists(t *testing.T) {
	env := make(map[string]interface{})
	env["path"] = "tmp"

	repo := Repo{FullName: "pearkes/test", SSHUrl: "foobaz"}
	env["repo"] = repo
	env["repo_state"] = "fetch"

	step := &StepFetchRepo{}

	results := step.Run(env)

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionHalt")
	}

	os.RemoveAll("tmp")
}

func TestStepFetchRepo_Skip_State(t *testing.T) {
	env := make(map[string]interface{})
	env["path"] = "tmp"

	repo := Repo{FullName: "pearkes/test", SSHUrl: "foobaz"}
	env["repo"] = repo
	env["repo_state"] = "clone"

	step := &StepFetchRepo{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
