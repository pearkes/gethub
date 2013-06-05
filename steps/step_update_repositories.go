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

	// Asynchronously update each repository
	var wg sync.WaitGroup
	for _, repo := range repos {
		wg.Add(1)
		go func(repo Repo) {
			switch checkRepo(repo, env) {
			case "fetch":
				fetches = append(fetches, repo.Name())
				fmt.Printf("%s.%s", green, clear)
			case "clone":
				clones = append(clones, repo.Name())
				fmt.Printf("%s.%s", green, clear)
			case "error":
				errors = append(errors, repo.Name())
				fmt.Printf("%s.%s", red, clear)
			case "ignore":
				ignores = append(ignores, repo.Name())
			}
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
