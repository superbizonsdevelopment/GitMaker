package manager

import(
  "fmt"
  "context"
  "github.com/google/go-github/github"
)

var (
  client = github.NewClient(nil)
)

func IsRepositoryExistOnOrganizationAccount(repoName string, repoAuthor string) (bool, string, string) {

  opt := &github.RepositoryListByOrgOptions{Type: "public"}

  repos, _, _ := client.Repositories.ListByOrg(context.Background(), repoAuthor, opt)

  for _, repo := range repos {
    if repoName == *repo.Name {
      return true, *repo.CloneURL, *repo.Language
    }
  }
  return false, "", ""
}

func IsRepositoryExistsOnUserAccount(repoName string, repoAuthor string) (bool, string, string) {
  client := github.NewClient(nil)
  stats, _, err := client.Admin.GetAdminStats(context.Background())
  if err != nil {
    return false, "", ""
  }
  fmt.Println(stats.GetRepos().TotalRepos)

  for _, repo := range stats.GetRepos() {
    fmt.Printf("Repo: %s", repo.Name())
  }
  return false, "", ""
}
