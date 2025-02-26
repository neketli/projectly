-- +goose Up
-- +goose StatementBegin
CREATE INDEX "idx_user_email" ON "users"("email");
CREATE INDEX "idx_task_project_id" ON "task"("status_id");
CREATE INDEX "idx_comment_task_id" ON "comment"("task_id");
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS "idx_user_email";
DROP INDEX IF EXISTS "idx_task_project_id";
DROP INDEX IF EXISTS "idx_comment_task_id";
-- +goose StatementEnd