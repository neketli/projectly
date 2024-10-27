-- +goose Up
-- +goose StatementBegin
CREATE TABLE "project" (
	"id" SERIAL PRIMARY KEY,
	"team_id" INTEGER NOT NULL,
	"title" VARCHAR(128) NOT NULL,
	"description" VARCHAR(128),
	"code" VARCHAR(5),
	FOREIGN KEY ("team_id") REFERENCES "team"("id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "project";
-- +goose StatementEnd