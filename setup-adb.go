package main

import "fmt"
import "os"
import "os/user"
import "path/filepath"
import "strings"
import "io"
import "net/http"

func expanduser(path string) string {

    if strings.HasPrefix(path, "~") == false {
        return path
    }

    usr, _ := user.Current()
    path = filepath.Join(usr.HomeDir, path[1:])
    return path
}

func mkdir(folder string) string {

    path := expanduser("~/atomic/" + folder)
    output := os.MkdirAll(path, os.ModePerm)
    if output == nil {
        fmt.Println("Folder created")
        return path
    } else {
        fmt.Println("Error: ", output)
        return "None"
    }
}

func download_file(filepath string, url string) (err error) {

  // Create the file
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()

  // Get the data, 
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func main() {
    args := os.Args[1:]

    path := mkdir(args[0])


    links := make(map[string]string)
    links["docker"] = "https://raw.githubusercontent.com/projectatomic/adb-atomic-developer-bundle/master/components/centos/centos-docker-base-setup/Vagrantfile"
    links["kubernetes"] = "https://raw.githubusercontent.com/projectatomic/adb-atomic-developer-bundle/master/components/centos/centos-k8s-singlenode-setup/Vagrantfile"
    links["openshift"] = "https://raw.githubusercontent.com/projectatomic/adb-atomic-developer-bundle/master/components/centos/centos-openshift-setup/Vagrantfile"
    links["installer"] = "https://raw.githubusercontent.com/surajssd/adb-post-installer/master/installer.sh"

    download_file(filepath.Join(path, "Vagrantfile"), links[args[1]])
    download_file(filepath.Join(path, "installer.sh"), links["installer"])
}
