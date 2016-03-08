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
vagrant ssh
```
