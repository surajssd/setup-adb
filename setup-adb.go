package main

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
    "io/ioutil"
    "github.com/surajssd/utils"
)

func mkdir(folder string) string {

    path := utils.ExpandUser("~/atomic/" + folder)
    output := os.MkdirAll(path, os.ModePerm)
    if output == nil {
        fmt.Println("Folder created")
        return path
    } else {
        fmt.Println("Error: ", output)
        return "None"
    }
}

func add_sshfs(vagrantfile string) {
    data, _ := ioutil.ReadFile(vagrantfile)

    lines := strings.Split(string(data), "\n")

    index := 0
    for i, line := range lines {
        if strings.Contains(line, "config.vm.network") {
            index = i
            break
        }
    }
    new_lines := append(append(lines[:index + 1], "  config.vm.synced_folder \"/home/hummer/git/atomic\", \"/home/vagrant/git\", type: \"sshfs\"\n"), lines[index+2:]...)

    output := strings.Join(new_lines, "\n")
    ioutil.WriteFile(vagrantfile, []byte(output), 0644)
}

func main() {
    args := os.Args[1:]

    path := mkdir(args[0])

    links := make(map[string]string)
    links["docker"] = "https://raw.githubusercontent.com/projectatomic/adb-atomic-developer-bundle/master/components/centos/centos-docker-base-setup/Vagrantfile"
    links["kubernetes"] = "https://raw.githubusercontent.com/projectatomic/adb-atomic-developer-bundle/master/components/centos/centos-k8s-singlenode-setup/Vagrantfile"
    links["openshift"] = "https://raw.githubusercontent.com/projectatomic/adb-atomic-developer-bundle/master/components/centos/centos-openshift-setup/Vagrantfile"
    links["installer"] = "https://raw.githubusercontent.com/surajssd/adb-post-installer/master/installer.sh"

    vagrantfile := filepath.Join(path, "Vagrantfile")
    utils.DownloadFile(vagrantfile, links[args[1]])
    utils.DownloadFile(filepath.Join(path, "installer.sh"), links["installer"])

    add_sshfs(vagrantfile)
}
