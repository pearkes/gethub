package steps

import (
	"os"
	"testing"

	"github.com/mitchellh/multistep"
)

func TestStepCheckRepo_Fetch(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.MkdirAll("tmp/pearkes/test/.git", 0777)
	env.Put("path", "tmp")
	env.Put("ignored_owners", []string{})
	env.Put("ignored_repos", []string{})

	repo := Repo{FullName: "pearkes/test"}

	env.Put("repo", repo)

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env.Get("repo_state").(string)

	if state != "fetch" {
		t.Fatal("repo state does not match fetch")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckRepo_Clone(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.MkdirAll("tmp", 0777)
	env.Put("path", "tmp")
	env.Put("ignored_owners", []string{})
	env.Put("ignored_repos", []string{})

	repo := Repo{FullName: "pearkes/test"}

	env.Put("repo", repo)

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env.Get("repo_state").(string)

	if state != "clone" {
		t.Fatal("repo state does not match clone")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckRepo_Ignore_Owner(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.MkdirAll("tmp", 0777)
	env.Put("path", "tmp")
	env.Put("ignored_owners", []string{"pearkes"})
	env.Put("ignored_repos", []string{})

	repo := Repo{FullName: "pearkes/test"}

	env.Put("repo", repo)

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env.Get("repo_state").(string)

	if state != "ignore" {
		t.Fatal("repo state does not match ignore")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckRepo_Ignore_Repo_Name(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.MkdirAll("tmp", 0777)
	env.Put("path", "tmp")
	env.Put("ignored_owners", []string{})
	env.Put("ignored_repos", []string{"bootstrap"})

	repo := Repo{FullName: "pearkes/bootstrap"}

	env.Put("repo", repo)

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env.Get("repo_state").(string)

	if state != "ignore" {
		t.Fatal("repo state does not match ignore")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}

func TestStepCheckRepo_Ignore_Repo_FullName(t *testing.T) {
	env := new(multistep.BasicStateBag)

	os.MkdirAll("tmp", 0777)
	env.Put("path", "tmp")
	env.Put("ignored_owners", []string{})
	env.Put("ignored_repos", []string{"pearkes/bootstrap"})

	repo := Repo{FullName: "pearkes/bootstrap"}

	env.Put("repo", repo)

	step := &StepCheckRepo{}

	results := step.Run(env)

	state := env.Get("repo_state").(string)

	if state != "ignore" {
		t.Fatal("repo state does not match ignore")
	}

	if results != multistep.ActionContinue {
		t.Fatal("step did not return ActionContinue")
	}

	os.RemoveAll("tmp")
}
