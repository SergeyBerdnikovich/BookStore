
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "auth_users_roles" ("user_id" integer,"role_id" integer, PRIMARY KEY ("user_id","role_id"));
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE auth_users_roles;
