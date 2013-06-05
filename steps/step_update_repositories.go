package steps

import (
	"fmt"
	"github.com/mitchellh/multistep"
	"log"
	"strings"
	"sync"
)

type StepUpdateRepositories struct{}

func (*StepUpdateRepositories) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Begin repository update sequence...")
	fmt.Printf("Contacting GitHub... ")
	repos := state["repos"].([]Repo)

	fmt.Printf("%sdone%s\n", GREEN, CLEAR)

	fmt.Printf("Updating repositories: ")

	fetches := []string{}
	clones := []string{}
	errors := []string{}

	steps := []multistep.Step{
		&StepCheckRepo{},
		&StepFetchRepo{},
		&StepCloneRepo{},
	}

	// Asynchronously update each repository
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo Repo) {
			state["repo"] = repo
			runner := &multistep.BasicRunner{Steps: steps}
			runner.Run(state)
			wg.Done()
		}(repo)
	}

	// Wait for every update to be finished
	wg.Wait()

	mess := []string{}

	if len(fetches) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d repos updated%s", GREEN, len(fetches), CLEAR))
	}

	if len(clones) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d new repos%s (%s)", GREEN, len(clones), CLEAR, strings.Join(clones, ", ")))
	}

	if len(errors) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d errors%s", RED, len(errors), CLEAR))
	}

	fmt.Printf("\n%s\n", strings.Join(mess, ", "))
	// Do some stuff
	return multistep.ActionContinue
}

func (*StepUpdateRepositories) Cleanup(map[string]interface{}) {}
