package util

import(
  "log"
  "io/ioutil"
  "../constants"
)

func GetClonedGitProjectPath(repoName string) string {
  files, err := ioutil.ReadDir(constants.ClonedAppsDir)
    if err != nil {
      log.Fatal(err)
    }

    for _, f := range files {
      if f.Name() == repoName && f.IsDir() {
        return constants.ClonedAppsDir + "/" + f.Name()
      }
    }

    return ""
}

func GetMainFileFromClonedGitProject(path string, repoName string) string {
  gitFiles, err := ioutil.ReadDir(path)
  if err != nil {
    log.Fatal(err)
  }

  for _, f := range gitFiles {
    if f.Name() == "main.go"  || f.Name() == repoName + ".go" || f.Name() == "server.go"{
      return f.Name()
    }
  }

  return ""
}
