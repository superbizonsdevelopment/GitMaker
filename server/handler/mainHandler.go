package handler

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"net/http"
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
			//If author is an organization

			client := github.NewClient(nil)

			opt := &github.RepositoryListByOrgOptions{Type: "public"}

			repos, _, _ := client.Repositories.ListByOrg(context.Background(), repoAuthor, opt)

			for _, repo := range repos {
				fmt.Printf("[DEBUG] Repo %s: %s\n", *repo.Owner.Login, *repo.ReleasesURL)
			}

		} else if repoIsAuthorIsAnOrganization == "no" {

		}
	}
}
