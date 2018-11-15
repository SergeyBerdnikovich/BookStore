
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

-- +swan safesql
ALTER TABLE audits RENAME TO audits_archive;
CREATE TABLE "audits" (
  "id" SERIAL,
  "created_at" timestamp with time zone,
  "operation_id" text,
  "resource" text,
  "resource_id" integer,
  "user_id" text,
  "field_name" text,
  "old_value" text,
  "new_value" text ,
  PRIMARY KEY ("id")
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE auditsV2;
