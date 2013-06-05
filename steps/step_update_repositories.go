package steps

import (
	"github.com/mitchellh/multistep"
)

type stepUpdateRepositories struct{}

func (*stepUpdateRepositories) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Begin repository update sequence...")
	fmt.Printf("Contacting GitHub... ")
	repos := listRemoteRepostories(env)

	fmt.Printf("%sdone%s\n", green, clear)

	fmt.Printf("Updating repositories: ")

	fetches := []string{}
	clones := []string{}
	errors := []string{}
	ignores := []string{}

	steps := []multistep.Step{
		&stepCheckRepo{},
		&stepCheckPath{},
		&stepFetchRepo{},
		&stepCloneRepo{},
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
		mess = append(mess, fmt.Sprintf("%s%d repos updated%s", green, len(fetches), clear))
	}

	if len(clones) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d new repos%s (%s)", green, len(clones), clear, strings.Join(clones, ", ")))
	}

	if len(errors) > 0 {
		mess = append(mess, fmt.Sprintf("%s%d errors%s", red, len(errors), clear))
	}

	fmt.Printf("\n%s\n", strings.Join(mess, ", "))
	// Do some stuff
	return multistep.ActionContinue
}

func (*stepUpdateRepositories) Cleanup(map[string]interface{}) {}
