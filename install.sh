#!/usr/bin/env bash
dep ensure -v
createdb -h localhost -p 5432 {|PROJECTNAME|}
swan -path=conf up
bee rs users-preload
createdb -h localhost -p 5432 {|PROJECTNAME|}_test
swan -path=conf -env=test up
