-- +goose Up
-- +goose StatementBegin
CREATE TABLE "team_user" (
	"team_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	FOREIGN KEY ("team_id") REFERENCES "team"("id") ON DELETE CASCADE,
	FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
	PRIMARY KEY ("team_id", "user_id")
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "team_user";
-- +goose StatementEnd