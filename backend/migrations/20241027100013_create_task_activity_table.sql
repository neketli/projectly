-- +goose Up
-- +goose StatementBegin
CREATE TABLE "task_activity" (
    "id" SERIAL PRIMARY KEY,
    "task_id" INTEGER NOT NULL REFERENCES "task"("id") ON DELETE CASCADE,
    "user_id" INTEGER NOT NULL REFERENCES "users"("id") ON DELETE CASCADE,
    "action_type" VARCHAR(32) NOT NULL,
    "field_name" VARCHAR(64),
    "old_value" TEXT,
    "new_value" TEXT,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX "idx_task_activity_task_id" ON "task_activity"("task_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "task_activity";
-- +goose StatementEnd
