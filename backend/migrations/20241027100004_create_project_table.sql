-- +goose Up
-- +goose StatementBegin
CREATE TABLE "project" (
	"id" SERIAL PRIMARY KEY,
	"title" VARCHAR(128) NOT NULL,
	"description" VARCHAR(128),
	"code" VARCHAR(5),
	"team_id" INTEGER NOT NULL,
	FOREIGN KEY ("team_id") REFERENCES "team"("id") ON DELETE CASCADE
);
ALTER TABLE "project"
ADD CONSTRAINT "project_code_unique" UNIQUE ("team_id", "code");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "project";
-- +goose StatementEnd