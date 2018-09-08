package handler

import (
	"fmt"
  "os/exec"
  "context"
  "net/http"
	"github.com/google/go-github/github"
)

func HandleMainFunction(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}

	if r.Method == http.MethodPost {
		repoName := r.PostFormValue("name")
		repoAuthor := r.PostFormValue("author")
		repoIsAuthorIsAnOrganization := r.PostFormValue("isOrganization")

		fmt.Fprintf(w, "Name: %v, Author: %v, isOrganization: %v\n", repoName, repoAuthor, repoIsAuthorIsAnOrganization)

		if repoIsAuthorIsAnOrganization == "yes" {
			//If author is an organnization

	     isRXOOA, repoLink := isRepositoryExistOnOrganizationAccount(repoName, repoAuthor)

       if isRXOOA == false {
         fmt.Fprintf(w, "Repository isn't exists!!!\n")
         return
       }

       cmd := exec.Command("./srcDownloader.sh", repoLink)
       cmd.Dir = "./scripts"

       out, err := cmd.Output()

       if err != nil {
         fmt.Println(err.Error())
         return
       }
       fmt.Println(repoLink)
       fmt.Println(string(out))

		} else if repoIsAuthorIsAnOrganization == "no" {

		}
	}
}

func isRepositoryExistOnOrganizationAccount(repoName string, repoAuthor string) (bool, string) {
  client := github.NewClient(nil)

  opt := &github.RepositoryListByOrgOptions{Type: "public"}

  repos, _, _ := client.Repositories.ListByOrg(context.Background(), repoAuthor, opt)

  for _, repo := range repos {
    if repoName == *repo.Name {
      return true, *repo.CloneURL
    }
  }
  return false, ""
}
