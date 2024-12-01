-- +goose Up
-- +goose StatementBegin
CREATE TABLE "roles" (
	"id" SERIAL PRIMARY KEY,
	"role_name" VARCHAR(128) NOT NULL
);
-- default roles
INSERT INTO "roles" ("role_name")
VALUES ('owner'),
	('editor'),
	('developer'),
	('user');
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "roles";
-- +goose StatementEnd