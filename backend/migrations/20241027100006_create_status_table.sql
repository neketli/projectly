-- +goose Up
-- +goose StatementBegin
CREATE TABLE "status" (
	"id" SERIAL PRIMARY KEY,
	"board_id" INTEGER NOT NULL,
	"title" VARCHAR(128) NOT NULL,
	"order" INTEGER NOT NULL,
	FOREIGN KEY ("board_id") REFERENCES "board"("id") ON DELETE CASCADE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "status";
-- +goose StatementEnd