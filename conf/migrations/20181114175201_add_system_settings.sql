
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "system_settings" ("id" serial,"created_at" timestamp with time zone,"updated_at" timestamp with time zone,"module" text NOT NULL,"key" text NOT NULL,"value" text NOT NULL,"description" text , PRIMARY KEY ("id"));
			  CREATE UNIQUE INDEX idx_module_key ON "system_settings"("module", "key");
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE system_settings;
