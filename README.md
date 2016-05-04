## Install

```bash
# get the shell file
curl -O https://raw.githubusercontent.com/surajssd/setup-adb/master/setup-adb
chmod a+x setup-adb
# put it in path
mkdir -p ~/.local/bin
mv setup-adb ~/.local/bin
```

## Usage

`setup-adb <folder-name> <provider-name>`

e.g.

```bash
setup-adb atomicappbug-45 openshift
cd ~/atomic/atomicappbug-45/
sed -i '/config.vm.network/a config.vm.synced_folder "/home/hummer/git/atomic", "/home/vagrant/git", type: "sshfs"' Vagrantfile 
vagrant up && vagrant ssh
```
