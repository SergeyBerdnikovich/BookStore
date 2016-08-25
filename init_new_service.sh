#!/usr/bin/env bash
PROJECTNAME=$1
echo "Building Project $PROJECTNAME"
# create project dir
echo "Cloning from skeleton..."
mkdir -p github.com/gtforge/$PROJECTNAME
cd github.com/gtforge/$PROJECTNAME
git clone git@github.com:gtforge/gett-skeleton-go.git .
brew install glide
glide install
go get bitbucket.org/liamstask/goose/cmd/goose
echo "Adding Flesh"
 $PROJECTNAME
sed -i -- 's/{{PROJECTNAME}}/$PROJECTNAME/g' .gitignore
sed -i -- 's/{{PROJECTNAME}}/$PROJECTNAME/g' Dockerfile
sed -i -- 's/{{PROJECTNAME}}/$PROJECTNAME/g' README.md
sed -i -- 's/{{PROJECTNAME}}/$PROJECTNAME/g' conf/dbconf.yml
echo "Tablizing..."
goose -env=development -path=config up
echo "Gitifying..."
rm -rf .git
git init
git add .
git commit -m "Initial Commit"
git branch dev
