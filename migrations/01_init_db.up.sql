CREATE TABLE IF NOT EXISTS "users" (
	"id" SERIAL PRIMARY KEY,
	"created_at" INTEGER,
	"updated_at" INTEGER,
	"deleted_at" INTEGER,
	"first_name" TEXT,
	"last_name" TEXT,
	"user_name" TEXT,
	"email" TEXT,
	"password" TEXT
);

CREATE TABLE IF NOT EXISTS "posts" (
	"id" SERIAL PRIMARY KEY,
	"created_at" INTEGER,
	"updated_at" INTEGER,
	"deleted_at" INTEGER,
	"title" TEXT,
	"body" TEXT,
	"user_id" INTEGER
);

-- ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id"):