
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- +swan safesql
ALTER TABLE "system_settings" ADD COLUMN "env" text NOT NULL DEFAULT '';
DROP INDEX idx_module_key;
CREATE UNIQUE INDEX idx_env_module_key ON "system_settings" ("env", "module", "key");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE env_column_to_system_settings;
