-- create "items" table
CREATE TABLE "items" ("id" uuid NOT NULL, "name" character varying NOT NULL, "status" character varying NOT NULL DEFAULT 'not_ready', "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, PRIMARY KEY ("id"));
