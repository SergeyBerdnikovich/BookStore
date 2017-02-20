
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "auth_sessions" ("session_key" text NOT NULL UNIQUE,"session_data" bytea,"session_expiry" timestamp with time zone );
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE auth_sessions;
