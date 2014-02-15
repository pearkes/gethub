package steps

import (
	"os"
	"os/exec"
	"testing"

	"github.com/mitchellh/multistep"
)

func TestStepCloneRepo_Exists(t *testing.T) {
	env := new(multistep.BasicStateBag)
	env.Put("path", "tmp")

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

	cmdInit.Run()
	cmdAdd.Run()
	cmdCommit.Run()

	repo := Repo{FullName: "pearkes/test", SSHUrl: "../origin"}
	env.Put("repo", repo)
	env.Put("repo_state", "clone")

	step := &StepCloneRepo{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCloneRepo_Not_Exists(t *testing.T) {
	env := new(multistep.BasicStateBag)
	env.Put("path", "tmp")

	repo := Repo{FullName: "pearkes/test", SSHUrl: "foobar"}
	env.Put("repo", repo)
	env.Put("repo_state", "clone")

	step := &StepCloneRepo{}

	results := step.Run(env)

	if results != multistep.ActionHalt {
		t.Fatal("step did not return ActionHalt")
	}

	os.RemoveAll("tmp")
}

func TestStepCloneRepo_Skip_State(t *testing.T) {
	env := new(multistep.BasicStateBag)
	env.Put("path", "tmp")

	originPath := "tmp/pearkes/origin"

	os.MkdirAll(originPath, 0777)

	repo := Repo{FullName: "pearkes/test", SSHUrl: "foobar"}
	env.Put("repo", repo)
	env.Put("repo_state", "fetch")

	step := &StepCloneRepo{}

	results := step.Run(env)

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
