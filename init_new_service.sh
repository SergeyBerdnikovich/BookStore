#!/usr/bin/env bash
brew update
brew install glide
go get github.com/gtforge/swan
go get github.com/beego/bee
go get github.com/onsi/ginkgo/ginkgo
echo "Adding Flesh"
LC_CTYPE=C sed -i '' 's/{|PROJECTNAME|}/'$PROJECTNAME'/g' $(find -L . -type f| grep -v .git | grep -v assets)
rm README.md
mv README_TEMPLATE.md README.md
echo "Gitifying..."
rm init_new_service.sh
git remote set-url origin git@github.com:gtforge/${PROJECTNAME}
git commit . -m "Cloned from skeleton project"
glide init --non-interactive
glide install
echo
echo "================================ DONE ================================== "
echo "The new project is located at \"${PWD}\" "
echo "Run install.sh to create a database, run migrations and seed users"
echo
