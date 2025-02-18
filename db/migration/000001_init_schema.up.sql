CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "nickname" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "chatrooms" (
  "id" serial PRIMARY KEY,
  "chatroom_name" varchar NOT NULL,
  "description" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "users_chatrooms" (
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL,
  "chatroom_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "role" varchar NOT NULL DEFAULT 'typer'
);

CREATE TABLE "messages" (
  "id" serial PRIMARY KEY,
  "user_chatroom_id" integer NOT NULL,
  "text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "chatrooms" ("chatroom_name");

CREATE UNIQUE INDEX ON "users_chatrooms" ("user_id", "chatroom_id");

CREATE INDEX ON "users_chatrooms" ("user_id");

CREATE INDEX ON "users_chatrooms" ("chatroom_id");

CREATE INDEX ON "messages" ("user_chatroom_id");

ALTER TABLE "users_chatrooms" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "users_chatrooms" ADD FOREIGN KEY ("chatroom_id") REFERENCES "chatrooms" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("user_chatroom_id") REFERENCES "users_chatrooms" ("id");
