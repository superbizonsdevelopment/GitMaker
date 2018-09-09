package manager

import (
  "fmt"
  "os/exec"
)

func UseBuildGoProjectScript(mainFile string) {
  cmd := exec.Command("./buildGoProjects.sh", mainFile)
  cmd.Dir = "./scripts"
  out, err := cmd.Output()

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println(string(out))
}

func UseSrcDownloaderScript(repoLink string) {
  cmd := exec.Command("./srcDownloader.sh", repoLink)
  cmd.Dir = "./scripts"

  out, err := cmd.Output()

  if err != nil {
    fmt.Println(err.Error())
    return
  }
  fmt.Println(repoLink)
  fmt.Println(string(out))
}
