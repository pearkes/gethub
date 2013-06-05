package steps

import (
	"github.com/mitchellh/multistep"
)

// type Repo represents a single repository
type Repo struct {
	FullName string `json:"full_name"`
	SSHUrl   string `json:"ssh_url"`
	HTTPSUrl string `json:"clone_url"`
}

func (r Repo) Owner() string {
	return strings.Split(r.FullName, "/")[0]
}

func (r Repo) Name() string {
	return strings.Split(r.FullName, "/")[1]
}

type stepCheckPath struct{}

func (*stepCheckPath) Run(state map[string]interface{}) multistep.StepAction {
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepCheckPath) Cleanup(map[string]interface{}) {}
