-- +goose Up
-- +goose StatementBegin
CREATE TABLE "task" (
	"id" SERIAL PRIMARY KEY,
	"project_index" INTEGER NOT NULL,
	"title" VARCHAR(128) NOT NULL,
	"description" VARCHAR(4096),
	"priority" INTEGER,
	"story_points" INTEGER,
	"tracked_time" INTEGER,
	"deadline" TIMESTAMP,
	"created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	"finished_at" TIMESTAMP,
	"status_id" INTEGER NOT NULL,
	"created_user_id" INTEGER NOT NULL,
	"assigned_user_id" INTEGER,
	FOREIGN KEY ("status_id") REFERENCES "status"("id") ON DELETE CASCADE,
	FOREIGN KEY ("created_user_id") REFERENCES "users"("id") ON DELETE CASCADE,
	FOREIGN KEY ("assigned_user_id") REFERENCES "users"("id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "task";
-- +goose StatementEnd