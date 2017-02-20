
-- NOTE: THIS FILE WAS PRODUCED BY THE Gett-Migrations(github.com/gtforge/services_common_go/gett-migrations)
-- DO NOT EDIT!!!!!!!
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "auth_users" ("id" serial,"created_at" timestamp with time zone,"updated_at" timestamp with time zone,"email" text NOT NULL UNIQUE,"encrypted_password" text,"reset_password_token" text,"reset_password_sent_at" timestamp with time zone,"remember_created_at" timestamp with time zone,"sign_in_count" integer,"current_sign_in_at" timestamp with time zone,"last_sign_in_at" timestamp with time zone,"current_sign_in_ip" text,"last_sign_in_ip" text,"encrypted_otp_secret" text,"encrypted_otp_secret_iv" text,"encrypted_otp_secret_salt" text,"consumed_timestep" integer , PRIMARY KEY ("id"));
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE auth_users;
