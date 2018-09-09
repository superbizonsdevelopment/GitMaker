package handler

import (
	"fmt"
	"../util"
  "context"
  "net/http"
	"../manager"
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

	     isRXOOA, repoLink, repoLanguage := isRepositoryExistOnOrganizationAccount(repoName, repoAuthor)

       if isRXOOA == false {
         fmt.Fprintf(w, "Repository isn't exists!!!\n")
         return
       }

       manager.UseSrcDownloaderScript(repoLink)

			 fmt.Println(util.GetClonedGitProjectPath(repoName))
			 fmt.Println(repoLanguage)

			 if  repoLanguage == "Go" {
			 		mainFile := util.GetMainFileFromClonedGitProject(util.GetClonedGitProjectPath(repoName), repoName)
					fmt.Printf("Main file: %v\n", mainFile)
					if mainFile == "" {
						fmt.Fprintf(w, "Can't find main file!\n")
						fmt.Println("Can't find main file!\n")
						return
					}

					manager.UseBuildGoProjectScript(mainFile, repoName)
			 } else if repoLanguage == "JavaScript" {
				 fmt.Fprintf(w, "JavaScript is can't compiled!\n")
				 return
			 }

		} else if repoIsAuthorIsAnOrganization == "no" {

		}
	}
}

func isRepositoryExistOnOrganizationAccount(repoName string, repoAuthor string) (bool, string, string) {
  client := github.NewClient(nil)

  opt := &github.RepositoryListByOrgOptions{Type: "public"}

  repos, _, _ := client.Repositories.ListByOrg(context.Background(), repoAuthor, opt)

  for _, repo := range repos {
    if repoName == *repo.Name {
      return true, *repo.CloneURL, *repo.Language
    }
  }
  return false, "", ""
}
