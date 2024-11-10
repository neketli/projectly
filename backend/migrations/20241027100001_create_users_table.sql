-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
	"id" SERIAL PRIMARY KEY,
	"name" VARCHAR(128) NOT NULL,
	"surname" VARCHAR(128) NOT NULL,
	"email" VARCHAR(128) UNIQUE NOT NULL,
	"password" VARCHAR(255) NOT NULL,
	"meta" JSON
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd