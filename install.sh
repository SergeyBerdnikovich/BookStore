#!/usr/bin/env bash
glide install
createdb {|PROJECTNAME|}
swan -path=conf up
