#!/usr/bin/env bash
brew update
brew install glide
go get github.com/gtforge/swan
go get github.com/beego/bee
go get github.com/onsi/ginkgo/ginkgo
echo "Adding Flesh"
LC_CTYPE=C sed -i '' 's/{|PROJECTNAME|}/'$PROJECTNAME'/g' $(find -L . -type f| grep -v .git | grep -v assets)
echo '# '$PROJECTNAME > README.md
chmod +x install.sh
echo "Gitifying..."
rm init_new_service.sh
rm -rf .git
glide install
git init
git add .
git commit -m "Initial Commit"
git branch dev
echo "Done! Run install.sh to create a database, run migrations and seed users"
