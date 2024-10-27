-- +goose Up
-- +goose StatementBegin
CREATE TABLE "attachment" (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(128) NOT NULL,
	"task_id" INTEGER NOT NULL,
	FOREIGN KEY ("task_id") REFERENCES "task"("id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "attachment";
-- +goose StatementEnd