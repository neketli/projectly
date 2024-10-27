-- +goose Up
-- +goose StatementBegin
CREATE TABLE "task" (
	"id" SERIAL PRIMARY KEY,
	"title" VARCHAR(128) NOT NULL,
	"description" VARCHAR(4096),
	"priority" INTEGER,
	"story_points" INTEGER,
	"created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updated_at" TIMESTAMP,
	"deadline" TIMESTAMP,
	"status_id" INTEGER NOT NULL,
	"created_user_id" INTEGER NOT NULL,
	"assigned_user_id" INTEGER,
	FOREIGN KEY ("status_id") REFERENCES "status"("id"),
	FOREIGN KEY ("created_user_id") REFERENCES "user"("id"),
	FOREIGN KEY ("assigned_user_id") REFERENCES "user"("id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "task";
-- +goose StatementEnd