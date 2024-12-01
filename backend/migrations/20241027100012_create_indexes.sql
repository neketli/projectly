-- +goose Up
-- +goose StatementBegin
CREATE INDEX "idx_user_email" ON "users"("email");
CREATE INDEX "idx_task_project_id" ON "task"("status_id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "idx_user_email";
DROP INDEX IF EXISTS "idx_task_project_id";
-- +goose StatementEnd