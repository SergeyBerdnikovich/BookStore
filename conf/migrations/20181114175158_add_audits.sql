
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE "audits" (
	"id" serial,
	"created_at" timestamp with time zone,
	"operation_id" varchar,
	"resource" varchar,
	"resource_id" integer,
	"user_id" integer,
	"field_name" varchar,
	"old_value" varchar,
	"new_value" varchar ,
	PRIMARY KEY ("id")
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE audits;
