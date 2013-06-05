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

type stepRetrieveRepositories struct{}

func (*stepRetrieveRepositories) Run(state map[string]interface{}) multistep.StepAction {
	log.Println("Retrieving remote repositories...")
	client := &http.Client{}

	req, err := http.NewRequest("GET",
		"https://api.github.com/user/repos?type=all&per_page=100", nil)

	req.SetBasicAuth(state["token"], "")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return multistep.ActionHalt
	}

	var repos []Repo

	err = json.Unmarshal(body, &repos)

	if err != nil {
		fmt.Println(err)
		return multistep.ActionHalt
	}

	log.Println(resp.Status)

	if resp.StatusCode != 200 {
		fmt.Println(red + "Uh oh, there was an error getting your repositories from GitHub. Here's what we got back:\n" + clear)
		fmt.Println(string(body))
		return multistep.ActionHalt
	}

	log.Println(len(repos), "repositories retrieved from GitHub")

	state["repos"] = repos

	return multistep.ActionContinue
}

func (*stepRetrieveRepositories) Cleanup(map[string]interface{}) {}
