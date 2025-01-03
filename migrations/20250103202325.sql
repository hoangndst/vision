-- Create "users" table
CREATE TABLE "public"."users" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "description" text NULL,
  "name" text NULL,
  "username" text NULL,
  "email" text NULL,
  "password" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- Create index "unique_user" to table: "users"
CREATE UNIQUE INDEX "unique_user" ON "public"."users" ("username", "email");
-- Create "blogs" table
CREATE TABLE "public"."blogs" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "description" text NULL,
  "created_by_id" uuid NULL,
  "path" text NULL,
  "raw_data" bytea NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_blogs_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_blogs_deleted_at" to table: "blogs"
CREATE INDEX "idx_blogs_deleted_at" ON "public"."blogs" ("deleted_at");
-- Create index "unique_blog" to table: "blogs"
CREATE UNIQUE INDEX "unique_blog" ON "public"."blogs" ("path");
-- Create "organizations" table
CREATE TABLE "public"."organizations" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "description" text NULL,
  "name" text NULL,
  "labels" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_organizations_deleted_at" to table: "organizations"
CREATE INDEX "idx_organizations_deleted_at" ON "public"."organizations" ("deleted_at");
-- Create index "unique_organization" to table: "organizations"
CREATE UNIQUE INDEX "unique_organization" ON "public"."organizations" ("name");
-- Create "project" table
CREATE TABLE "public"."project" (
  "id" uuid NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "description" text NULL,
  "created_by_id" uuid NULL,
  "name" text NULL,
  "organization_id" uuid NULL,
  "path" text NULL,
  "labels" text NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_project_created_by" FOREIGN KEY ("created_by_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_project_organization" FOREIGN KEY ("organization_id") REFERENCES "public"."organizations" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_project_deleted_at" to table: "project"
CREATE INDEX "idx_project_deleted_at" ON "public"."project" ("deleted_at");
-- Create index "unique_project" to table: "project"
CREATE UNIQUE INDEX "unique_project" ON "public"."project" ("name", "path");
