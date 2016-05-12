## Install

Setting up `go` environment, if you already have one skip these steps
```bash
# install go
sudo dnf -y install go

# if you already have one skip this step
mkdir $HOME/go

export GOPATH=$HOME/go
echo 'export GOPATH=$HOME/go' >> ~/.bash_profile

export PATH=$PATH:$GOPATH/bin
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bash_profile
```

Install
```bash
go get github.com/surajssd/setup-adb
```

## Usage

`setup-adb <folder-name> <provider-name>`

e.g.

```bash
setup-adb atomicappbug-45 openshift
cd ~/atomic/atomicappbug-45/
vagrant up && vagrant ssh
```
