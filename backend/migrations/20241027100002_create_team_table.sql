-- +goose Up
-- +goose StatementBegin
CREATE TABLE "team" (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(128) NOT NULL,
	"description" VARCHAR(256)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "team";
-- +goose StatementEnd