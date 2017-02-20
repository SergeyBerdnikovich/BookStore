
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "auth_roles" ("id" serial,"name" text,"resource_id" integer,"resource_type" text,"created_at" timestamp with time zone,"updated_at" timestamp with time zone , PRIMARY KEY ("id"));
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE auth_roles;
