-- +goose Up
-- +goose StatementBegin
CREATE TABLE "board" (
	"id" SERIAL PRIMARY KEY,
	"project_id" INTEGER NOT NULL,
	"title" VARCHAR(128) NOT NULL,
	FOREIGN KEY ("project_id") REFERENCES "project"("id") ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "board";
-- +goose StatementEnd