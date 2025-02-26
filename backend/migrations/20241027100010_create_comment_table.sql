-- +goose Up
-- +goose StatementBegin
CREATE TABLE "comment" (
	"id" SERIAL PRIMARY KEY,
	"text" VARCHAR(1024) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	"updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	"task_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	FOREIGN KEY ("task_id") REFERENCES "task"("id") ON DELETE CASCADE,
	FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "comment";
-- +goose StatementEnd