package steps

import (
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/mitchellh/multistep"
)

type StepUpdateRepositories struct{}

func (*StepUpdateRepositories) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Begin repository update sequence...")

	repos := state.Get("repos").([]Repo)

	fmt.Printf("Updating repositories: ")

	fetches := []string{}
	clones := []string{}
	errors := []string{}
	ignores := []string{}

	maxConcurrent := 32
	var sem = make(chan int, maxConcurrent) // Counting semaphore

	// Asynchronously update each repository
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)

		go func(repo Repo) {
			sem <- 1 // Wait

			// New state for each repo update runner
			cState := new(multistep.BasicStateBag)

			steps := []multistep.Step{
				&StepCheckRepo{},
				&StepFetchRepo{},
				&StepCloneRepo{},
			}

			// Copy parent state values over.
			cState.Put("ignored_repos", state.Get("ignored_repos"))
			cState.Put("ignored_owners", state.Get("ignored_owners"))
			cState.Put("repo_state", state.Get("repo_state"))
			cState.Put("repo_result", state.Get("repo_result"))
			cState.Put("path", state.Get("path"))

			cState.Put("repo", repo)

			runner := &multistep.BasicRunner{Steps: steps}

			runner.Run(cState)

			switch cState.Get("repo_result").(string) {
			case "fetch":
				fetches = append(fetches, repo.Name())
			case "clone":
				clones = append(clones, repo.Name())
			case "error":
				errors = append(errors, repo.Name())
			case "ignore":
				ignores = append(ignores, repo.Name())
			}
			<-sem // Signal
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
		mess = append(mess, fmt.Sprintf("%s%d errors%s (%s)", RED, len(errors), CLEAR, strings.Join(errors, ", ")))
	}

	log.Println("Ignored repositories: ", strings.Join(ignores, ", "))

	fmt.Printf("\n%s\n", strings.Join(mess, ", "))
	// Do some stuff
	return multistep.ActionContinue
}

func (*StepUpdateRepositories) Cleanup(multistep.StateBag) {}
