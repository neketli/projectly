-- +goose Up
-- +goose StatementBegin
CREATE TABLE "team_roles" (
	"team_id" INTEGER NOT NULL,
	"user_id" INTEGER NOT NULL,
	"role_id" INTEGER NOT NULL,
	FOREIGN KEY ("team_id") REFERENCES "team"("id") ON DELETE CASCADE,
	FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE,
	FOREIGN KEY ("role_id") REFERENCES "roles"("id") ON DELETE CASCADE,
	PRIMARY KEY ("team_id", "user_id", "role_id")
);
ALTER TABLE "team_roles"
ADD CONSTRAINT "team_user_unique" UNIQUE ("team_id", "user_id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "team_roles";
-- +goose StatementEnd