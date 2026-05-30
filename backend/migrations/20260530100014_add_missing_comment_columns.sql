-- +goose Up
-- +goose StatementBegin
ALTER TABLE "comment"
	ADD COLUMN "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
	ADD COLUMN "user_id" INTEGER NOT NULL DEFAULT 0;

ALTER TABLE "comment"
	ADD FOREIGN KEY ("user_id") REFERENCES "users"("id") ON DELETE CASCADE;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE "comment"
	DROP CONSTRAINT IF EXISTS "comment_user_id_fkey",
	DROP COLUMN "user_id",
	DROP COLUMN "updated_at";
-- +goose StatementEnd
