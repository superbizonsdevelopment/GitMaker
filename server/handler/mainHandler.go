package handler

import (
	"fmt"
	"../util"
  "net/http"
	"../manager"
	"../constants"
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

	     isREOOA, repoLink, repoLanguage := manager.IsRepositoryExistOnOrganizationAccount(repoName, repoAuthor)

       if isREOOA == false {
         fmt.Fprintf(w, "Repository isn't exists!!!\n")
         return
       }

       manager.UseSrcDownloaderScript(repoLink)

			 fmt.Println(util.GetClonedGitProjectPath(repoName))
			 fmt.Println(repoLanguage)

			 if  repoLanguage == "Go" {
			 		mainFile := util.GetMainFileFromClonedGitProject(util.GetClonedGitProjectPath(repoName), repoName)

					if mainFile == "" {
						fmt.Fprintf(w, "Can't find main file!\n")
						fmt.Println("Can't find main file!\n")
						return
					}

					manager.UseBuildGoProjectScript(mainFile, repoName)

					sendFile(w, r, repoName)

					manager.UseUnnecesserySrcRemover(repoName)
			 } else if repoLanguage == "JavaScript" {
				 fmt.Fprintf(w, "JavaScript can't compiled!\n")
				 return
			 }

		} else if repoIsAuthorIsAnOrganization == "no" {
			isREOUA, repoLink, repoLanguage := manager.IsRepositoryExistsOnUserAccount(repoName, repoAuthor)

			if isREOUA == false {
				fmt.Fprintf(w, "Repository isn't exists!!!\n")
				return
			}
		}
	}
}

func sendFile(w http.ResponseWriter, r *http.Request, repoName string) {
	fmt.Println(constants.ClonedReposDir + "/" + repoName)
  http.ServeFile(w, r, constants.ClonedReposDir + "/" + repoName)
}
