package handler

import (
	"fmt"
  "context"
  "net/http"
	"github.com/google/go-github/github"
)

var (
  repoLink = ""
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
         //send packet then repo isn't exist!
         return
       }

       fmt.Println(repoLink)

       //switch on sh program with link to repo argument

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
      return true, *repo.URL
    }
  }
  return false, ""
}
