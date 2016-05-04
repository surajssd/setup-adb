#!/bin/bash

mkdir -p ~/git/atomic

cd ~/git/atomic
git clone https://github.com/surajssd/atomicapp
cd atomicapp/
git remote add upstream https://github.com/projectatomic/atomicapp.git
git pull --ff upstream master

cat > .git/hooks/pre-push <<EOF
#!/bin/bash
flake8 -v atomicapp
EOF

cd ~/git/atomic
git clone https://github.com/surajssd/openshift2nulecule.git
cd openshift2nulecule/
git remote add upstream https://github.com/projectatomic/openshift2nulecule.git
git pull --ff upstream master


cd ~/git/atomic
git clone https://github.com/surajssd/nulecule-library.git
cd nulecule-library/
git remote add upstream https://github.com/projectatomic/nulecule-library.git
git pull --ff upstream master

