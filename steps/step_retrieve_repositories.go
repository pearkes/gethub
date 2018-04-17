package steps

import (
	"context"
	"log"
	"strings"

	"github.com/google/go-github/github"
	"github.com/mitchellh/multistep"
	"golang.org/x/oauth2"
)

// type Org represents an organization
type Org struct {
	Name string `json:"login"`
}

// type Repo represents a single repository
type Repo struct {
	FullName string
	SSHUrl   string
	HTTPSUrl string
}

func (r Repo) Owner() string {
	return strings.Split(r.FullName, "/")[0]
}

func (r Repo) Name() string {
	return strings.Split(r.FullName, "/")[1]
}

type StepRetrieveRepositories struct{}

func (*StepRetrieveRepositories) Run(state multistep.StateBag) multistep.StepAction {
	log.Println("Retrieving remote repositories...")
	var allRepos []Repo
	token := state.Get("token").(string)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	log.Printf("Contacting GitHub... ")

	// Retrieve Organizations
	orgs, _, err := client.Organizations.List(ctx, "", nil)

	if err != nil {
		log.Println(err)
	}
	reposForOrgs := map[*github.Organization]*[]*github.Repository{}
	for _, org := range orgs {
		log.Println("Getting", org)
		repos, err := reposForOrg(ctx, client, *org.Login)
		if err != nil {
			log.Println("Could not get repos for ", org.GetName(), err)
			continue
		} else {
			log.Printf("Got %d repos for %s", len(repos), org.GetName())
		}
		for _, repo := range repos {
			allRepos = append(allRepos, gRepo(repo))
		}
		reposForOrgs[org] = &repos
	}

	memberOpts := &github.RepositoryListOptions{
		Type: "member",
	}

	ownerOpts := &github.RepositoryListOptions{
		Type: "owner",
	}

	memeberRepos, err := getAllRepos(ctx, client, memberOpts)
	if err != nil {
		log.Printf("Error getting repos for Owner: %s\n", err)
	}

	ownerRepos, err := getAllRepos(ctx, client, ownerOpts)
	if err != nil {
		log.Printf("Error getting repos for Owner: %s\n", err)
	}
	for _, r := range memeberRepos {
		allRepos = append(allRepos, gRepo(r))
	}
	for _, r := range ownerRepos {
		allRepos = append(allRepos, gRepo(r))
	}

	log.Println(len(allRepos), "repositories retrieved from GitHub")
	state.Put("repos", allRepos)
	log.Printf("%sdone%s\n", GREEN, CLEAR)

	return multistep.ActionContinue
}

func gRepo(repo *github.Repository) Repo {
	log.Println("Got:" + repo.GetFullName())
	return Repo{
		FullName: repo.GetFullName(),
		SSHUrl:   repo.GetSSHURL(),
		HTTPSUrl: repo.GetCloneURL(),
	}
}
func getAllRepos(ctx context.Context, client *github.Client, opts *github.RepositoryListOptions) ([]*github.Repository, error) {
	// get all pages of results
	var allRepos []*github.Repository

	for {
		log.Println("Fetching repos ")
		repos, resp, err := client.Repositories.List(ctx, "", opts)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return allRepos, nil
}

func reposForOrg(ctx context.Context, client *github.Client, org string) ([]*github.Repository, error) {
	perPage := 100
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: perPage},
	}

	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return nil, err
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	log.Printf("Got %d repos for %s\n", len(allRepos), org)
	return allRepos, nil
}

func (*StepRetrieveRepositories) Cleanup(multistep.StateBag) {}
