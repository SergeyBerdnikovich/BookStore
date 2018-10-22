#!/usr/bin/env bash
go get -u github.com/golang/dep/cmd/dep
go get -u github.com/gtforge/swan
go get -u github.com/beego/bee
go get -u github.com/onsi/ginkgo/ginkgo
go get -u github.com/gtforge/services_common_go/gett-auth/seed
echo "Adding Flesh"
LC_CTYPE=C sed -i '' 's/{|PROJECTNAME|}/'$PROJECTNAME'/g' $(find -L . -type f| grep -v .git | grep -v assets && echo .gitignore)
rm README.md
mv README_TEMPLATE.md README.md
echo "Gitifying..."
rm init_new_service.sh
git remote set-url origin git@github.com:gtforge/${PROJECTNAME}
git commit . -m "Cloned from skeleton project"
dep ensure -vendor-only
echo
echo "================================ DONE ================================== "
echo "The new project is located at \"${PWD}\" "
echo "Run install.sh to create a database, run migrations and seed users"
echo
