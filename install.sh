#!/usr/bin/env bash
glide install
createdb permission_enforcement_service
swan -path=conf up
bee rs users-preload
createdb permission_enforcement_service_test
swan -path=conf -env=test up
