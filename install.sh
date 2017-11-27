#!/usr/bin/env bash
dep ensure -v --update
createdb {|PROJECTNAME|}
swan -path=conf up
bee rs users-preload
createdb {|PROJECTNAME|}_test
swan -path=conf -env=test up
