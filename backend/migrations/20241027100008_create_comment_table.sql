-- +goose Up
-- +goose StatementBegin
CREATE TABLE "comment" (
	"id" SERIAL PRIMARY KEY,
	"text" VARCHAR(1024) NOT NULL,
	"created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	"task_id" INTEGER NOT NULL,
	FOREIGN KEY ("task_id") REFERENCES "task"("id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "comment";
-- +goose StatementEnd