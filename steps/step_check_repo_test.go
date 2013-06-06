package steps

import (
	"github.com/mitchellh/multistep"
	"os"
	"testing"
)

func TestStepCheckRepo_Fetch(t *testing.T) {
	env := make(map[string]interface{})

	os.MkdirAll("tmp/pearkes/test/.git", 0777)
	env["path"] = "tmp"
	env["ignored_owners"] = []string{}
	env["ignored_repos"] = []string{}

	repo := Repo{FullName: "pearkes/test"}

	env["repo"] = repo

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env["repo_state"].(string)

	if state != "fetch" {
		t.Fatal("repo state does not match fetch")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckRepo_Clone(t *testing.T) {
	env := make(map[string]interface{})

	os.MkdirAll("tmp", 0777)
	env["path"] = "tmp"
	env["ignored_owners"] = []string{}
	env["ignored_repos"] = []string{}

	repo := Repo{FullName: "pearkes/test"}

	env["repo"] = repo

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env["repo_state"].(string)

	if state != "clone" {
		t.Fatal("repo state does not match clone")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckRepo_Ignore_Owner(t *testing.T) {
	env := make(map[string]interface{})

	os.MkdirAll("tmp", 0777)
	env["path"] = "tmp"
	env["ignored_owners"] = []string{"pearkes"}
	env["ignored_repos"] = []string{}

	repo := Repo{FullName: "pearkes/test"}

	env["repo"] = repo

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env["repo_state"].(string)

	if state != "ignore" {
		t.Fatal("repo state does not match ignore")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
